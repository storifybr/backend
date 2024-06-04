package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/storify/backend/internal/schemas"
)

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

	companyResponse := schemas.CompanyResponse {
		ID: company.ID,
		Name: company.Name,
		TaxID: company.TaxID,
		Email: company.Email,
		Phone: company.Phone,
		CreatedAt: company.CreatedAt,
		UpdatedAt: company.UpdatedAt,
		DeletedAt: company.DeletedAt,
	}

	sendSuccess(ctx, "company creation", companyResponse)
}
