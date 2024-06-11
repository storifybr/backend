package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/storify/backend/internal/schemas"
)

// @BasePath /api/v1

// @Summary Delete a company
// @Description Delete a company
// @Tags Companies
// @Accept json
// @Produce json
// @Param id query string true "Company identifier"
// @Success 200 {object} CompanySuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /company [delete]
func DeleteCompanyHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		err := errParamIsRequired("id", "queryParameter")
		logger.Errorf(err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	company := schemas.Company{}

	if err := db.First(&company, id).Error; err != nil {
		logger.Errorf("error getting company: %v", err)
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("company with id %s not found", id))
		return
	}

	if err := db.Delete(&company).Error; err != nil {
		logger.Errorf("error deleting company: %v", err)
		sendError(ctx, http.StatusBadRequest, fmt.Sprintf("failed to delete company with id %s", id))
		return
	}

	sendSuccess(ctx, "company deleted", company)
}
