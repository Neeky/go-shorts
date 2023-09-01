package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
)

func main() {
	// var db *sql.DB
	// var err error

	// 连接配置
	var config = mysql.Config{
		User:   "appuser",
		Passwd: "123456",
		Net:    "tcp",
		Addr:   "127.0.0.1:3308",
		DBName: "tempdb",
	}

	// 验证参数
	db, err := sql.Open("mysql", config.FormatDSN())
	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(8)
	db.SetConnMaxLifetime(time.Second * 3600 * 4)
	defer db.Close()

	if err != nil {
		fmt.Printf("connect to database fail %v", err)
	}

	// 连接数据库看能不能成功
	pingerr := db.Ping()
	if pingerr != nil {
		fmt.Printf("ping fail %v \n", pingerr)
	} else {
		fmt.Printf("connect to %s succ \n", config.Addr)
	}

	time.Sleep(time.Second * 7)
}
