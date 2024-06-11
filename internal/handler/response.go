package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/storify/backend/internal/schemas"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"message": msg,
	})
}

func sendSuccess(ctx *gin.Context, operation string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s successfully", operation),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type CompanySuccessResponse struct {
	Message string                  `json:"message"`
	Data    schemas.CompanyResponse `json:"data"`
}
