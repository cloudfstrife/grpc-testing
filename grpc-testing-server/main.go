package main

import (
	"app/grpc-testing-server/api/echo"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type EchoServer struct {
	echo.UnimplementedEchoServer
}

func (e *EchoServer) Echo(ctx context.Context, par *echo.EchoRequest) (*echo.EchoResponse, error) {
	result := &echo.EchoResponse{
		Msg: par.GetMsg(),
	}
	return result, nil
}

func main() {
	log.Println("Go")
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	echo.RegisterEchoServer(s, &EchoServer{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
