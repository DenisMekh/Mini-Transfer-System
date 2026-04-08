package main

import (
	"net"

	"github.com/DenisMekh/mini-transfer-system/account-svc/internal/handler"
	pb "github.com/DenisMekh/mini-transfer-system/gen/go/account"
	"google.golang.org/grpc"
)

type Server struct {
	grpcServer *grpc.Server
	listener   net.Listener
}

func New(handler *handler.AccountHandler, addr string) (*Server, error) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	srv := grpc.NewServer()
	pb.RegisterAccountServiceServer(srv, handler)
	return &Server{
		grpcServer: srv,
		listener:   lis,
	}, nil
}

func (s *Server) Run() error {
	return s.grpcServer.Serve(s.listener)
}

func (s *Server) Stop() {
	s.grpcServer.GracefulStop()
}
