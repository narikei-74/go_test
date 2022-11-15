package main

import (
	"github.com/gin-gonic/gin"
	"github.com/narikei-74/go_test/controller"
)

func main() {
  r := gin.Default()

  // test1
  r.POST("/test1", controller.test1)
  // test2
  r.POST("/test2", controller.test2)

  r.Run(":8080")
}