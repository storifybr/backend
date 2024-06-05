package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/storify/backend/internal/schemas"
)

func GetCompaniesHandler(ctx *gin.Context) {
	companies := []schemas.Company{}

	if err := db.Find(&companies).Error; err != nil {
		logger.Errorf("error getting companies: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error getting companies")
		return
	}

	sendSuccess(ctx, "companies retrieved", companies)
}
