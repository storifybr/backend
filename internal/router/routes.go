package router

import (
	"github.com/gin-gonic/gin"
	"github.com/storify/backend/internal/handler"

	docs "github.com/storify/backend/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{
		v1.GET("/companies", handler.GetCompaniesHandler)
		v1.GET("/company", handler.GetCompanyHandler)
		v1.POST("/company", handler.CreateCompanyHandler)
		v1.PUT("/company", handler.UpdateCompanyHandler)
		v1.DELETE("/company", handler.DeleteCompanyHandler)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
