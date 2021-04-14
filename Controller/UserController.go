package Controller

import (
	"ginweb02/Server"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

// 这里是构造函数
func NewUserController() *UserController {
	return &UserController{}
}

// 这里是业务方法
func (this *UserController) GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"data": "hello world2",
		})
	}
}

// 这里是处理路由的地儿
func (this *UserController) Router(server *Server.Server) {
	server.G.Handle("GET", "/", this.GetUser())
}
