package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "cloud_native/cloudstorage/transferFileService/pbfile"
)

var grpcAddrs = "localhost:8000"

func main() {
	grpcServer := grpc.NewServer()
	pb.RegisterFileTransferServiceServer(grpcServer, &pb.Server{})
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