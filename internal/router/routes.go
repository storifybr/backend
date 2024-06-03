package router

import (
	"github.com/gin-gonic/gin"
	"github.com/storify/backend/internal/handler"
)

func initializeRoutes(r *gin.Engine) {
	handler.InitializeHandler()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/companies", handler.GetCompaniesHandler)
		v1.GET("/company", handler.GetCompanyHandler)
		v1.POST("/company", handler.CreateCompanyHandler)
		v1.PUT("/company", handler.UpdateCompanyHandler)
		v1.DELETE("/company", handler.DeleteCompanyHandler)
	}
}
