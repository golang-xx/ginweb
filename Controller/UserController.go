package Controller

import (
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
func (this *UserController) RouterGroup(g *gin.RouterGroup) {
	g.Handle("GET", "/", this.GetUser())
}
