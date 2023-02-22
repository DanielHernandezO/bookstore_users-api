package config

import "github.com/gin-gonic/gin"

type Strategy interface {
	Start()
}

func NewStrategy() Strategy {
	return &webAPI{}
}

type webAPI struct{}

func (*webAPI) Start() {
	router := configApplication()
	if err := router.Run(":8080"); err != nil {
		//handle the error --> to do
		return
	}
}

func configApplication() *gin.Engine {
	router := gin.Default()
	mapUrls(router, buildDependencies(router))
	return router
}
