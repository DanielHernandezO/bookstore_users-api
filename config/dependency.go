package config

import (
	"github.com/DanielHernandezO/bookstore_users-api/controller"
	"github.com/gin-gonic/gin"
)

type dependencies struct {
	pingController controller.PingController
}

func buildDependencies(router *gin.Engine) *dependencies {
	return &dependencies{
		pingController: controller.NewPingController(),
	}
}
