package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/storify/backend/internal/schemas"
)

// @BasePath /api/v1

// @Summary Update a company
// @Description Update a company
// @Tags Companies
// @Accept json
// @Produce json
// @Param id query string true "Company identifier"
// @Param request body UpdateCompanyRequest true "Company data to update"
// @Success 200 {object} UpdateCompanyResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /company [put]
func UpdateCompanyHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		err := errParamIsRequired("id", "queryParameter")
		logger.Errorf(err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	request := UpdateCompanyRequest{}

	if err := validateJsonRequest(ctx, &request); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	company := schemas.Company{}

	if err := db.Find(&company, id).Error; err != nil {
		logger.Errorf("error getting company: %v", err.Error())
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("company with id %s not found", id))
		return
	}

	if request.Name != "" {
		company.Name = request.Name
	}

	if request.TaxID != "" {
		company.TaxID = request.TaxID
	}

	if request.Email != "" {
		company.Email = request.Email
	}

	if request.Phone != "" {
		company.Phone = request.Phone
	}

	if err := db.Save(&company).Error; err != nil {
		logger.Errorf("error updating company: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("failed to update company with id %s", id))
		return
	}

	sendSuccess(ctx, "company updated", company)
}
