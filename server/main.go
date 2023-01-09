package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/tony-zhuo/grpc-practice/protos/user"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) GetUserInfo(ctx context.Context, req *pb.UserInfoReq) (*pb.UserInfoResp, error) {
	return &pb.UserInfoResp{LastName: "Zhuo", FirstName: "Tony", Age: int32(28)}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("server net listen error: %+v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("server Serve err: %+v", err)
	}
}
