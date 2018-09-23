package server

import (
	"net"

	pb "github.com/mrbeskin/array-sender/grpc/data"
	"google.golang.org/grpc"
)

type Server struct {
	Port string
}

func (s *Server) ListenGrpc() error {
	lis, err := net.Listen("tcp", s.Port)
	if err != nil {
		return fmt.Errorf("unable to start tcp server: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterDataListener(grpcServer, s)
	pb.RegisterDataWriter(grpcServer, s)
	grpcServer.Serve(lis)
}
