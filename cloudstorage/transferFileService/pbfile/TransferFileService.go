package pbfile

import (
	"fmt"
	"os"
	"io"
	"io/ioutil"
	"context"
)

type Server struct {
	UnimplementedFileTransferServiceServer
}

func (s *Server) DownloadFile(req *DownloadRequest, stream FileTransferService_DownloadFileServer) error {
	content := make([]byte, 1024)
	file, err := os.Open("./file/" + req.Username + "/" + req.FileName)
	if err != nil {
		fmt.Println("download open file error:", err)
		return err
	}
	defer file.Close()
	for {
		num, err := file.Read(content)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		stream.Send(&DownloadResponse{
			FileContent : content[:num],
		})
	}
	return nil
}

func (s *Server) UploadFile(stream FileTransferService_UploadFileServer) error {
	var filename string
	var username string
	var content []byte
	var err error

	for {
		req, err := stream.Recv()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("receive request error:", err)
			return err
		}
		filename = req.FileName
		username = req.Username
		content = append(content, req.FileContent...)
	}

	var res *UploadResponse
	if err != nil {
		res = &UploadResponse{
			Success : false,
			Message : filename + " upload failed!",
		}
	} else {
		res = &UploadResponse{
			Success : true,
			Message : filename + " upload success!",
		}
	}

	err = stream.SendAndClose(res)
	if err != nil {
		fmt.Println("SendAndClose error:", err)
		return err
	}
	dest := "./file/" + username + "/" + filename
	fmt.Println("username:", username)
	f, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("OpenFile error:", err)
		return err
	}
	defer f.Close()
	_, err = f.Write(content)
	if err != nil {
		fmt.Println("write into file error:", err)
		return err
	}
	return nil
}

func(s *Server) ShowDir(ctx context.Context, req *UploadRequest) (*DirResponse, error) {
	fileInfo, err := ioutil.ReadDir("./file/" + req.Username)
	if err != nil {
		fmt.Println(req.Username, " read dir error:",  err)
		return &DirResponse{}, err
	}

	var filenames []string
	for _, k := range fileInfo {
		filenames = append(filenames, k.Name())
	}

	fmt.Println("filenames:", filenames)

	return &DirResponse{
		Contents : filenames,
	}, nil
}

func(s *Server) CreateDir(ctx context.Context, req *UploadRequest) (*UploadRequest, error) {
	err := os.Mkdir("./file/" + req.Username, 0666)
	if err != nil {
		fmt.Println("mkdir error:", err)
		return &UploadRequest{}, err
	}

	return &UploadRequest{}, nil
}

func(s *Server) DeleteFile(ctx context.Context, req *UploadRequest) (*UploadRequest, error) {
	os.Remove("./file/" + req.Username + "/" + req.FileName)
	// if err != nil {
	// 	fmt.Println("remove file error:", err)
	// 	return nil, err
	// }

	return &UploadRequest{}, nil
}