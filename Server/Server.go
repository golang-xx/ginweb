package Server

import (
	"github.com/gin-gonic/gin"
)

// 这里是定义一个接口，解决上述弊端的规范性
type IController interface {
	// 这个传参就是脚手架主程
	Router(server *Server)
}

// 定义一个脚手架
type Server struct {
	*gin.Engine
	// 路由分组一会儿会用到
	g *gin.RouterGroup
}

// 初始化函数
func Init() *Server {
	// 作为Server的构造器
	s := &Server{Engine: gin.New()}
	// 返回作为链式调用
	return s
}

// 监听函数，更好的做法是这里的端口应该放到配置文件
func (this *Server) Listen() {
	this.Run(":8092")
}

// 这里是路由的关键代码，这里会挂载路由
func (this *Server) Route(controllers ...IController) *Server {
	// 遍历所有的控制层，这里使用接口，就是为了将Router实例化
	for _, c := range controllers {
		c.Router(this)
	}
	return this
}
