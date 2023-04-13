package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "cloudstorage/login/pbfile"
)

var grpcAddrs = "localhost:8001"

func main() {
	grpcServer := grpc.NewServer()
	pb.RegisterLoginServiceServer(grpcServer, &pb.Server{})
	listener, err := net.Listen("tcp", grpcAddrs)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("start listen to " + grpcAddrs)

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalln(err)
	}
}