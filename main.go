package main
import "github.com/gin-gonic/gin"

func main(){
	r := gin.Default()
	r.GET("/",func(c *gin.Context){
		c.JSON(200,gin.H{"message":"pong1",})
	})
	r.Run(":8092")
}