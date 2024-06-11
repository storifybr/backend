package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/storify/backend/internal/schemas"
)

// @BasePath /api/v1

// @Summary Create a company
// @Description Create a new company
// @Tags Companies
// @Accept json
// @Produce json
// @Param request body CreateCompanyRequest true "Company data"
// @Success 200 {object} CreateCompanyResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /company [post]
func CreateCompanyHandler(ctx *gin.Context) {
	request := CreateCompanyRequest{}

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

	company := schemas.Company {
		Name: request.Name,
		TaxID: request.TaxID,
		Email: request.Email,
		Phone: request.Phone,
	}

	if err := db.Create(&company).Error; err != nil {
		logger.Errorf("error creating company: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating company")
		return
	}

	sendSuccess(ctx, "company created", company)
}
