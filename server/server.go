package server

import (
	"bufio"
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"net"

	pb "github.com/mrbeskin/intds/grpc/data"
	"google.golang.org/grpc"
)

// TODO: Document this, or handle arbitrarily large payloads - the goal is for all TCP traffic to only require a single packet every time.
const MAX_MSG_SIZE = 4096 // NOTE: The size of the array is bounded by this - receiving payloads of larger than this size will cause corruption

// TcpServer is the internal server used to send and receive messages with a client.
type TcpServer struct {
	Port string
	out  chan pb.IntDataArray
	in   chan pb.IntDataArray
}

func NewTcpServer(port string) *TcpServer {
	return &TcpServer{
		Port: port,
		out:  make(chan pb.IntDataArray),
		in:   make(chan pb.IntDataArray),
	}
}

// ListenTcp() Starts the Tcp which will listen for a connection, then use that connection to read Int data arrays and pass them off to channels
// for the Grpc consumers to handle.
func (s *TcpServer) ListenTcp() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", s.Port))
	if err != nil {
		panic(fmt.Errorf("unable to start tcp server for client/server: %v", err))
	}
	fmt.Println(fmt.Printf("listening on: %v\n", lis.Addr()))
	conn, err := lis.Accept()
	if err != nil {
		panic(fmt.Errorf("unable to accept connection from client: %v", err))
	}
	buf := make([]byte, MAX_MSG_SIZE)
	reader := bufio.NewReader(conn)
	go s.handleSends(conn)
	for {
		fmt.Println("connected to client")
		size, err := reader.Read(buf)
		if err != nil {
			fmt.Println(fmt.Errorf("failure reading from connection: %v", err))
		}
		ida := pb.IntDataArray{}
		err = binary.Read(bytes.NewBuffer(buf[:size]), binary.BigEndian, &ida)
		if err != nil {
			fmt.Println(fmt.Errorf("failed to convert binary data to data array: %v", err))
		}
		s.out <- ida
	}
}

// handleSends  takes IntDataArrays to the client.
func (s *TcpServer) handleSends(conn net.Conn) {
	writer := bufio.NewWriter(conn)
	for data := range s.in {
		buf := &bytes.Buffer{}
		binary.Write(buf, binary.BigEndian, data)
		_, err := writer.Write(buf.Bytes())
		if err != nil {
			fmt.Println(fmt.Errorf("could not send data: %v", err))
		}
	}
}

// GrpcServer is the server which listens for local clients and handles GrpcRequests
type GrpcServer struct {
	tcps *TcpServer
}

// GetIntDataArrays implements the gRPC function which streams data to the gRPC client.
func (grpcs *GrpcServer) GetIntDataArrays(in *pb.GetIntDataArrayStreamRequest, stream pb.DataServer_GetIntDataArraysServer) error {
	for d := range grpcs.tcps.out {
		err := stream.Send(&d)
		if err != nil {
			return err
		}
	}
	return nil
}

// WriteIntDataArray implements the gRPCfunction that allows the server to write a data array to the client
func (grpcs *GrpcServer) WriteIntDataArray(context context.Context, dataArray *pb.IntDataArray) (*pb.ServerSendResponse, error) {
	// TODO: Implement using handle sends
	grpcs.tcps.in <- *dataArray
	return &pb.ServerSendResponse{}, nil
}

// ListenGrpc starts the server and its internal tcp server and begins accepting connections from clients
func (grpcs *GrpcServer) ListenGrpc() error {
	lis, err := net.Listen("tcp", "0")
	if err != nil {
		return fmt.Errorf("unable to start tcp server for grpc: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterDataServerServer(grpcServer, grpcs)
	grpcServer.Serve(lis)
	return nil
}
