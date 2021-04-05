package http

import "github.com/gin-gonic/gin"

type Response struct {
	ctx  *gin.Context
	code int
}

func (r *Response) json(data interface{}) {

}
