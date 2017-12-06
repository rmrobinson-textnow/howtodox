package main

import (
	"fmt"
	"log"
	"net"

	"github.com/rmrobinson-textnow/howtodox/golang/contexts/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"time"
)

type srv struct {
}

func (s *srv) TestFunction(ctx context.Context, req *proto.TestRequest) (*proto.TestResponse, error) {
	if req.A == "slow" {
		time.Sleep(time.Second * 10)

		return &proto.TestResponse{
			B: "slow",
		}, nil
	} else {
		return &proto.TestResponse{
			B: "fast",
		}, nil
	}
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 10101))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	s := &srv{}
	proto.RegisterTestServer(grpcServer, s)
	grpcServer.Serve(lis)
}
