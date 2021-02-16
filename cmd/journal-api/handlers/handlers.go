package handlers

import (
	"github.com/gin-gonic/gin"
)

func Mainhandler(c *gin.Context) {
	c.JSON(200, "Hello world")
}
