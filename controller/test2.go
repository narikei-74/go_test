package controller

import (
	"net/http"

  "github.com/gin-gonic/gin"
)

func test2(c *gin.Context) {
  c.String(http.StatusOK, "test2")
}