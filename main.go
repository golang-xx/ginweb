package main

import (
	"ginweb02/Controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	Controller.NewUserController(r).Router()
	r.Run(":8092")
}
