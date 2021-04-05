package user

import (
	"github.com/alex-1900/sparkling/app/service"
	"github.com/gin-gonic/gin"
)

func actionRegister(c *gin.Context) {
	serviceAuth := service.GetAuth()
	c.JSON(200, gin.H{
		"message": serviceAuth.Register(),
	})
}
