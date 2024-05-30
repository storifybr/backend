package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCompanyHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET Company",
	})
}
