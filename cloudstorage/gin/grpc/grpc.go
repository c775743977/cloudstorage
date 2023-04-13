package grpc

import (
	"log"
	_"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "cloudstorage/gin/pbfile"
)

var loginAddrs = "localhost:8001"
var uploadAddrs = "localhost:8000"

var LoginClient pb.LoginServiceClient
var UploadClient pb.FileTransferServiceClient


func init()  {
	loginconn, err := grpc.Dial(loginAddrs, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}

	LoginClient = pb.NewLoginServiceClient(loginconn)

	uploadconn, err := grpc.Dial(uploadAddrs, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}

	UploadClient = pb.NewFileTransferServiceClient(uploadconn)
}