package main

import (
	"app/grpc-testing-client/api/echo"
	"context"
	"fmt"
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
	r, err := c.Echo(ctx)
	if err != nil {
		log.Fatalf("echo failed :%#v", err)
	}
	for i := 0; i < 10; i++ {
		err = r.Send(&echo.EchoRequest{Msg: fmt.Sprintf("hello - %d", i)})
		if err != nil {
			log.Fatal("send msg error", err)
		}

		resp, err := r.Recv()
		if err != nil {
			log.Fatal("recv msg error", err)
		}
		fmt.Println(resp.Msg)
	}
	err = r.CloseSend()
	if err != nil {
		log.Fatal("close send error", err)
	}
}
