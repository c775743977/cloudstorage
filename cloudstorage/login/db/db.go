package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/redis/go-redis/v9"
)

var mysqlAddrs = "root:Chen@123@tcp(localhost:3306)/cloudstorage"


var RDB *redis.ClusterClient
var MDB *gorm.DB
var err error

func init() {
	MDB, err = gorm.Open(mysql.Open(mysqlAddrs), &gorm.Config{})
	if err != nil {
		fmt.Println("connect to mysql error:", err)
	}
	RDB = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs : []string{"192.168.108.165:6381", "192.168.108.165:6382", "192.168.108.165:6383"},
	})
}