package server

import (
	"net"

	pb "github.com/mrbeskin/array-sender/grpc/data"
	"google.golang.org/grpc"
)

// TODO: Document this, or handle arbitrarily large payloads - the goal is for all TCP traffic to only require a single packet every time.
const MAX_BUF_SIZE = 4096 // NOTE: The size of the array is bounded by this - receiving payloads of larger than this size will cause corruption

// TcpServer is the internal server used to send and receive messages with a client.
type TcpServer struct {
	Port int
	out  <-chan pb.IntDataArray
	in   chan<- pb.IntDataArray
}

// ListenTcp() Starts the Tcp which will listen for a connection, then use that connection to read Int data arrays and pass them off to channels
// for the Grpc consumers to handle.
func (s *TcpServer) ListenTcp() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		return fmt.Errorf("unable to start tcp server for client/server: %v", err)
	}
	for {
		connection, err := lis.Accept()
		if err != nil {
			return fmt.Errorf("unable to accept connection from client: %v", err)
		}
		size, _ := strconv.ParseInt(strings.Trim(string(bufferFileSize), ":"), 10, 64)
	}
}

// GrpcServer is the server which listens for local clients and handles GrpcRequests
type GrpcServer struct {
	tcps *TcpServer
}

// GetIntDataArrays implements the gRPC function which streams data to the gRPC client.
func (grpcs *GrpcServer) GetIntDataArrays(ctx context.Context, in *pb.GetIntDataArrayStreamRequest, stream pb.Data_GetIntDataArraysServer) error {
	for d := range grpcs.tcps.out {
		err := stream.Send(d)
		if err != nil {
			return err
		}
	}
	return nil
}

// WriteIntDataArray implements the gRPCfunction that allows the server to write a data array to the client
func (grpcs *GrpcServer) WriteIntDataArray(dataArray *pb.IntDataArray) (*pb.ServerSendResponse, error) {
	s.GetClientConnection()
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
}
