package config

import "github.com/gin-gonic/gin"

func mapUrls(router *gin.Engine, dedependencies *dependencies) {
	router.GET("/ping", dedependencies.pingController.Ping)
}
