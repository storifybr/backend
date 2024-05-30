package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteCompanyHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "DELETE Company",
	})
}
