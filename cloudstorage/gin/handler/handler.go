package handler

import (
	"fmt"
	"context"
	"io"

	rpc "cloudstorage/gin/grpc"
	pb "cloudstorage/gin/pbfile"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	ctx := context.Background()
	cookie, _ := c.Cookie("uuid")
	if cookie == "" {
		c.HTML(200, "index.html", nil)
		return
	}

	res, err := rpc.LoginClient.GetSession(ctx, &pb.Cookie{
		Code : cookie,
	})
	if err != nil {
		c.String(500, "服务器内部错误")
		return
	}

	fileinfo, err := rpc.UploadClient.ShowDir(ctx, &pb.UploadRequest{
		Username : res.Username,
	})

	res.Contents = fileinfo.Contents
	c.HTML(200, "index.html", res)
}

func LoginHandler(c *gin.Context) {
	ctx := context.Background()
	var user pb.LoginRequest
	err := c.Bind(&user)
	if err != nil {
		c.HTML(400, "login.html", "无法提交空数据")
	}

	res, err := rpc.LoginClient.Login(ctx, &user)
	if err != nil {
		c.String(500, "服务器内部错误")
		return
	}

	if res.Username == "" {
		c.HTML(400, "login.html", "用户名或密码不正确")
		return
	}

	c.SetCookie("uuid", res.Cookie, 0, "/", "localhost", false, true)
	c.HTML(200, "login_success.html", res.Username)
}

func RegistHandler(c *gin.Context) {
	cookie, _ := c.Cookie("uuid")
	if cookie != "" {
		c.HTML(200, "verify.html", nil)
		return
	}
	ctx := context.Background()
	var user pb.RegistInfo
	err := c.Bind(&user)
	if err != nil {
		c.HTML(400, "regist.html", "无法提交空数据")
		return
	}
	if user.Password != user.Repwd {
		c.HTML(400, "regist.html", "两次输入密码不一致")
		return
	}

	res, err := rpc.LoginClient.Regist(ctx, &pb.RegistRequest{
		Username : user.Username,
		Password : user.Password,
		Email : user.Email,
	})
	if err != nil {
		c.String(500, "服务器内部错误")
		return
	}

	rpc.UploadClient.CreateDir(ctx, &pb.UploadRequest{
		Username : user.Username,
	})

	c.SetCookie("uuid", res.Cookie, 300, "/", "localhost", false, true)
	c.HTML(200, "verify.html", nil)
}

func VerifyHandler(c *gin.Context) {
	ctx := context.Background()
	code := c.PostForm("code")
	cookie, err := c.Cookie("uuid")
	if err != nil {
		c.String(400, "发生错误,请重新注册")
	}

	res, err := rpc.LoginClient.Verify(ctx, &pb.VerifyRequest{
		Cookie : cookie,
		Code : code,
	})
	if err != nil {
		c.String(500, "服务器内部错误")
		return
	}

	if res.Result == false {
		c.HTML(400, "verify.html", "验证码错误")
		return
	}

	c.HTML(200, "regist_success.html", nil)
}

func LogoutHandler(c *gin.Context) {
	ctx := context.Background()
	cookie, _ := c.Cookie("uuid")
	if cookie == "" {
		c.String(400, "当前并未处登录状态")
		return
	}

	c.SetCookie("uuid", "", -1, "/", "localhost", false, true)
	_, err := rpc.LoginClient.Logout(ctx, &pb.Cookie{
		Code : cookie,
	})
	if err != nil {
		c.String(500, "服务器内部错误")
		return
	}
	c.HTML(200, "index.html", nil)
}

func CheckLogin(c *gin.Context) *pb.LoginResponse {
	ctx := context.Background()
	cookie, _ := c.Cookie("uuid")
	if cookie == "" {
		return nil
	}

	res, err := rpc.LoginClient.GetSession(ctx, &pb.Cookie{
		Code : cookie,
	})
	if err != nil {
		c.String(500, "服务器内部错误")
		return nil
	}
	return res
}

func UploadHandler(c *gin.Context) {
	ctx := context.Background()
	user := CheckLogin(c)
	if user == nil {
		return
	}

	fileheader, err := c.FormFile("file")
	fmt.Println("file-size:", fileheader.Size)
	if fileheader.Size > 1024 * 1024 * 100 {
		if user.VIP != "Y" {
			c.HTML(400, "oversize.html", nil)
			return
		}
	}
	if err != nil {
		fmt.Println("get file error:", err)
		c.HTML(400, "index.html", pb.LoginRequest{
			Username : user.Username,
		})
		return
	}

	file, err := fileheader.Open()
	if err != nil {
		fmt.Println("open file error:", err)
		c.HTML(400, "index.html", pb.LoginRequest{
			Username : user.Username,
		})
		return
	}
	//每次上传1kb的数据, 可以通过这个切片长度来限制上传大小
	var content = make([]byte, 1024)
	stream, err := rpc.UploadClient.UploadFile(ctx)
	if err != nil {
		fmt.Println(err)
		c.String(500, "服务器内部错误")
		return
	}

	for {
		num, err := file.Read(content)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("client stream read file error:", err)
			c.String(500, "服务器内部错误")
			return
		}

		err = stream.Send(&pb.UploadRequest{
			Username : user.Username,
			FileName : fileheader.Filename,
			FileContent : content[:num],
		})
		if err != nil {
			fmt.Println("client stream send byte[] error:", err)
			c.String(500, "服务器内部错误")
			return
		}
	}

	res, err := stream.CloseAndRecv()
	fmt.Println("res:", res)
	if err != nil {
		c.String(500, fmt.Sprint(err))
	} else {
		c.HTML(200, "upload_success.html", fileheader.Filename)
	}
}

func DownloadHandler(c *gin.Context) {
	user := CheckLogin(c)
	if user == nil {
		return
	}

	filename := c.Query("filename")
	// filename := c.Request.FormValue("filename")
	
	//可以设置切片长度来限制下载大小
	var content []byte

	stream, err := rpc.UploadClient.DownloadFile(context.Background(), &pb.DownloadRequest{
		Username : user.Username,
		FileName : filename,
	})
	if err != nil {
		c.String(500, "服务器内部错误")
		return
	}

	for {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("stream receive error:", err)
			return
		}
		content = append(content, res.FileContent...)
	}

	c.Header("Content-Disposition", "attachment; filename=" + filename)
	c.Header("Content-Type", "application/octet-stream")

	_, err = c.Writer.Write(content)
	if err != nil {
		fmt.Println("write into response error:", err)
		return
	}

	fmt.Println("用户:", user.Username, "\t开始下载文件:", filename)

	IndexHandler(c)
}

func RegistVIPHandler(c *gin.Context) {
	ctx := context.Background()
	var user pb.RegistInfo
	c.Bind(&user)
	ruser := CheckLogin(c)
	if ruser == nil {
		c.String(400, "尚未登录")
		return
	}

	if ruser.Username != user.Username {
		c.String(400, "登录信息错误")
		return
	}
	res, _ := rpc.LoginClient.RegistVIP(ctx, &pb.Cookie{
		Code : ruser.Cookie,
	})
	
	if res.Result {
		c.SetCookie("uuid", "", -1, "/", "localhost", false, true)
		_, err := rpc.LoginClient.Logout(ctx, &pb.Cookie{
			Code : ruser.Cookie,
		})
		if err != nil {
			c.String(500, "服务器内部错误")
			return
		}
		c.HTML(200, "registVIP_success.html", nil)
		return
	} else {
		c.String(400, "开通失败")
		return
	}
}

func DeleteFileHandler(c *gin.Context) {
	ctx := context.Background()
	filename := c.Query("filename")
	user := CheckLogin(c)
	fmt.Println("filename:", filename)
	_, err := rpc.UploadClient.DeleteFile(ctx, &pb.UploadRequest{
		Username : user.Username,
		FileName : filename,
	})
	if err != nil {
		c.String(500, "服务器内部错误")
		return
	}

	IndexHandler(c)
}