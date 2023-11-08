package main

import (
	"app/grpc-testing-client/api/echo"
	"context"
	"errors"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := echo.NewEchoClient(conn)
	ctx, cancle := context.WithCancel(context.Background())
	defer cancle()
	r, err := c.Echo(ctx, &echo.EchoRequest{Msg: "hello"})
	if err != nil {
		log.Fatalf("echo failed :%#v", err)
	}
	for {
		resp, err := r.Recv()
		if err == nil {
			fmt.Println(resp.Msg)
			continue
		}

		if errors.Is(err, io.EOF) {
			break
		}
		log.Fatal("error", err)
	}
}
