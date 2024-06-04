package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func validateJsonRequest(ctx *gin.Context, request interface{}) error {
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return fmt.Errorf("malformed request body")
	}

	return nil
}

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateCompanyRequest struct {
	Name  string `json:"name"`
	TaxID string `json:"taxId"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (r *CreateCompanyRequest) Validate() error {
	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}

	if r.TaxID == "" {
		return errParamIsRequired("taxId", "string")
	}

	if r.Email == "" {
		return errParamIsRequired("email", "string")
	}

	if r.Phone == "" {
		return errParamIsRequired("phone", "string")
	}

	return nil
}
