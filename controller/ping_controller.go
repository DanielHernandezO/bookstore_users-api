package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingController interface {
	Ping(c *gin.Context)
}

type pingController struct{}

func NewPingController() *pingController {
	return &pingController{}
}

func (*pingController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "Pong")
}
