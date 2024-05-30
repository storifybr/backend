package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateCompanyHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "PUT Company",
	})
}
