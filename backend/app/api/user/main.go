package user

import "github.com/gin-gonic/gin"

func Handle(r *gin.RouterGroup) {
	auth := r.Group("user")
	auth.POST("register", actionRegister)
}
