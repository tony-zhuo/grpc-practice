package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/tony-zhuo/grpc-practice/protos/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	userInfo, err := c.GetUserInfo(ctx, &pb.UserInfoReq{Account: "0928144268"})
	if err != nil {
		log.Fatalf("grpc GetUserInfo err: %+v", err)
	}
	fmt.Printf("grpc GetUserInfo: %+v", userInfo)
}
