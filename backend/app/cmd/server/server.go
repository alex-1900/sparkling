package server

import (
	"github.com/alex-1900/sparkling/app/api/user"
	"github.com/gin-gonic/gin"
)

func startServer(addr ...string) {
	r := gin.Default()
	root := r.Group("")
	user.Handle(root.Group("user"))
	r.Run(addr...)
}
