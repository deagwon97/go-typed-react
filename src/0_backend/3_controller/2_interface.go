package controller

import "github.com/gin-gonic/gin"

type ControllerInterface interface {
	CreateFile(c *gin.Context)
	Ping(c *gin.Context)
}
