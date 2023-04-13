package pbfile

import (
	"log"
	"context"
	"fmt"
	"errors"
	"time"

	"cloudstorage/login/db"
	"cloudstorage/login/utils"
)

type Server struct {
	UnimplementedLoginServiceServer
}

func(s *Server) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	var user LoginRequest
	res := db.MDB.Where("username = ?", req.Username).Find(&user)
	if res.Error != nil {
		log.Fatalln("db find error:", res.Error)
	}

	if req.Password != user.Password {
		return &LoginResponse{
			Username : "",
		}, nil
	}

	cookie := utils.CreateUUID()

	err := db.RDB.HMSet(ctx , cookie, "username", req.Username, "password", req.Password, "email", req.Email, "VIP", user.VIP).Err()
	if err != nil {
		fmt.Println("redis hmset error", err)
		return &LoginResponse{}, err
	}

	return &LoginResponse{
		Username : req.Username,
		Cookie : cookie,
	}, nil
}

func(s *Server) Regist(ctx context.Context, req *RegistRequest) (*RegistResponse, error) {
	if utils.CheckExist(req.Username) {
		return &RegistResponse{}, errors.New("该用户名已存在")
	}

	cookie := utils.CreateUUID()
	code := utils.CreateCode()
	fmt.Println("cookie:", cookie)
	fmt.Println("reqinfo:", req)
	err := db.RDB.HMSet(ctx , cookie, "username", req.Username, "password", req.Password, "email", req.Email, "code", code).Err()
	if err != nil {
		fmt.Println("redis hmset error", err)
		return &RegistResponse{}, err
	}

	db.RDB.Expire(ctx, cookie, time.Minute * 3)

	utils.SendMail(code, req.Email)

	return &RegistResponse{
		Cookie : cookie,
	}, nil
}

func(s *Server) Verify(ctx context.Context, req *VerifyRequest) (*VerifyResponse, error) {
	data, err := db.RDB.HMGet(context.Background(), req.Cookie, "code", "username", "password", "email").Result()
	if err != nil {
		fmt.Println("RDB.HMGet error:", err)
		return nil, err
	}
	if req.Code != data[0].(string) {
		return &VerifyResponse{
			Result : false,
		}, nil
	}

	var user = LoginRequest{
		Username : data[1].(string),
		Password : data[2].(string),
		Email : data[3].(string),
	}
	res := db.MDB.Create(&user)
	if res.Error != nil {
		fmt.Println("save user to mysql error:", res.Error)
		return nil, res.Error
	}

	return &VerifyResponse{
		Result : true,
	}, nil
}

func(s *Server) GetSession(ctx context.Context, req *Cookie) (*LoginResponse, error) {
	data, err := db.RDB.HMGet(context.Background(), req.Code, "username", "VIP").Result()
	if err != nil {
		fmt.Println("RDB.HMGet error:", err)
		return nil, err
	}

	return &LoginResponse{
		Username : data[0].(string),
		VIP : data[1].(string),
		Cookie : req.Code,
	}, nil
}

func(s *Server) Logout(ctx context.Context, req *Cookie) (*Cookie, error) {
	err := db.RDB.HDel(ctx, req.Code, "username", "password", "email").Err()
	if err != nil {
		fmt.Println("redis delete session error:", err)
		return nil, err
	}
	return &Cookie{}, nil
}

func(s *Server) RegistVIP(ctx context.Context, req *Cookie) (*VerifyResponse, error) {
	data, err := db.RDB.HMGet(context.Background(), req.Code, "username").Result()
	if err != nil {
		fmt.Println("RDB.HMGet error:", err)
		return nil, err
	}

	username := data[0].(string)
	res := db.MDB.Model(&LoginResponse{}).Where("username = ?", username).Update("VIP", "Y")
	if res.Error != nil {
		fmt.Println("registVIP error:", res.Error)
		return &VerifyResponse{
			Result : false,
		}, nil
	}

	return &VerifyResponse{
		Result : true,
	}, nil
}