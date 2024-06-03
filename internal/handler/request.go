package handler

import (
	"fmt"
)

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
