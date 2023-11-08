package main

import (
	"app/grpc-testing-server/api/echo"
	"log"
	"net"

	"google.golang.org/grpc"
)

type EchoServer struct {
	echo.UnimplementedEchoServer
}

func (e *EchoServer) Echo(par *echo.EchoRequest, echoserver echo.Echo_EchoServer) error {
	msg := par.GetMsg()
	for i := 0; i < 10; i++ {
		echoserver.Send(&echo.EchoResponse{Msg: msg})
	}
	return nil
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
