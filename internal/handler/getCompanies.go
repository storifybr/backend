package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/storify/backend/internal/schemas"
)

// @BasePath /api/v1

// @Summary Get companies
// @Description Get all companies
// @Tags Companies
// @Accept json
// @Produce json
// @Success 200 {object} GetCompaniesResponse
// @Failure 500 {object} ErrorResponse
// @Router /companies [get]
func GetCompaniesHandler(ctx *gin.Context) {
	companies := []schemas.Company{}

	if err := db.Find(&companies).Error; err != nil {
		logger.Errorf("error getting companies: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error getting companies")
		return
	}

	sendSuccess(ctx, "companies retrieved", companies)
}
