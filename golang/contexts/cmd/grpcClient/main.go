package main

import (
	"fmt"
	"time"

	"github.com/rmrobinson-textnow/howtodox/golang/contexts/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial("127.0.0.1:10101", opts...)

	if err != nil {
		fmt.Printf("Error connecting: %s\n", err.Error())
		return
	}

	client := proto.NewTestClient(conn)

	req := &proto.TestRequest{
		A: "asdf",
	}

	fmt.Printf("Making test request\n")

	ctx, _ := context.WithTimeout(context.Background(), time.Second)

	resp, err := client.TestFunction(ctx, req)

	if err != nil {
		fmt.Printf("Error running: %s\n", err.Error())
		return
	}

	fmt.Printf("%+v\n", resp)
}
