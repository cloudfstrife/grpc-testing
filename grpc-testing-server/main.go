package main

import (
	"app/grpc-testing-server/api/echo"
	"errors"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type EchoServer struct {
	echo.UnimplementedEchoServer
}

func (e *EchoServer) Echo(server echo.Echo_EchoServer) error {
	for {
		r, err := server.Recv()
		if err == nil {
			server.Send(&echo.EchoResponse{Msg: "return " + r.Msg})
			continue
		}
		if errors.Is(err, io.EOF) {
			break
		}
		return err
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
