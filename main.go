package main

import (
	"./controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	NewUserController(r).Router()
	r.Run(":8092")
}
