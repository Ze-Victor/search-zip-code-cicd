package main

import (
	docs "github.com/Ze-Victor/search-zip-code/docs"
	auth "github.com/Ze-Victor/search-zip-code/internal/pkg/authorization"
	cep "github.com/Ze-Victor/search-zip-code/internal/pkg/cep"
	health "github.com/Ze-Victor/search-zip-code/internal/pkg/health"
	metrics "github.com/Ze-Victor/search-zip-code/internal/pkg/metrics"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	apiV1 := router.Group(basePath)

	//Routes CEP
	apiV1.GET("/cep/:cep", auth.AuthMiddleware(), cep.SearchCEPHandler)

	//Routes Auth
	apiV1.POST("/auth", auth.CreateTokenHandler)

	//Routes Health
	apiV1.GET("/health", health.CheckApplicationHealth)

	//Routes Docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//Routes Metrics
	metrics.RegisterMetrics(router)

}
