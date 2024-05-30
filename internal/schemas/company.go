package schemas

import (
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	ID    int    `json:"id"`
	Name  string `json:"name"`
	TaxID string `json:"taxId"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
