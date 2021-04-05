package main

import (
	"github.com/alex-1900/sparkling/app/cmd/db"
	"github.com/alex-1900/sparkling/app/cmd/root"
	"github.com/alex-1900/sparkling/app/cmd/server"
)

func main() {
	cmd := root.New()
	// 服务器相关命令
	server.New("server").Handle(cmd)
	// 数据库相关命令
	db.New("db").Handle(cmd)
	cmd.GetCmd().Execute()
}
