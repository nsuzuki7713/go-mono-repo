package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// 1. 8080番portのリスナーを作成
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	// 2. gRPCサーバを作成
	s := grpc.NewServer()

	// 3. 作成したサーバーを8080番ポートで稼働させる
	go func() {
		log.Printf("start grpc server port: %v", port)
		s.Serve(listener)
	}()
}
