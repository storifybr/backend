package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCompaniesHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET Companies",
	})
}

func GetCompanyHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET Company",
	})
}

func CreateCompanyHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "POST Company",
	})
}

func UpdateCompanyHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "PUT Company",
	})
}

func DeleteCompanyHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "DELETE Company",
	})
}
