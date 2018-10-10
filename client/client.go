package client

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

type TcpClient struct {
	serverHost string
	serverPort string
	in         chan pb.IntDataArray
	out        chan pb.IntDataArray
}

func NewTcpClient(host string, port string) *TcpClient {
	return &TcpClient{
		serverHost: host,
		serverPort: port,
		in:         make(chan pb.IntDataArray),
	}
}

func (tcpClient *TcpClient) DialTcp() {
	conn, err := net.Dial("tcp", fmt.Sprintf("%v:%v", tcpClient.serverHost, tcpClient.serverPort))
	if err != nil {
		panic(err)
	}
	go tcpClient.handleIncoming(conn)
	tcpClient.handleOutgoing(conn)
}

const MAX_MSG_SIZE = 4096 // TODO: config

func (tcpClient *TcpClient) handleIncoming(conn net.Conn) {
	reader := bufio.NewReader(conn)
	buf := make([]byte, MAX_MSG_SIZE)
	for {
		size, err := reader.Read(buf)
		ida := pb.IntDataArray{}
		err = binary.Read(bytes.NewBuffer(buf[:size]), binary.BigEndian, &ida)
		if err != nil {
			panic(fmt.Errorf("failed to convert binary data to data array: %v", err))
		}
		tcpClient.in <- ida
	}
}

func (tcpClient *TcpClient) handleOutgoing(conn net.Conn) {
	writer := bufio.NewWriter(conn)
	for data := range tcpClient.out {
		buf := &bytes.Buffer{}
		binary.Write(buf, binary.BigEndian, data)
		_, err := writer.Write(buf.Bytes())
		if err != nil {
			fmt.Println(fmt.Errorf("could not send data: %v", err))
		}
	}
}

func (tcpClient *TcpClient) WriteIntDataArray(ctx context.Context, dataArray *pb.IntDataArray) (*pb.IntDataArray, error) {
	tcpClient.out <- *dataArray
	arrayIn := <-tcpClient.in
	return &arrayIn, nil
}

func (tcpClient *TcpClient) ListenGrpc() error {
	lis, err := net.Listen("tcp", "0")
	if err != nil {
		return fmt.Errorf("unable to start tcp server for grpc: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterDataClientServer(grpcServer, tcpClient)
	grpcServer.Serve(lis)
	return nil
}
