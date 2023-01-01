package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) CreateFile(c *gin.Context) {
	c.JSON(http.StatusOK, "hello!!")
}

func (ctrl *Controller) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}
