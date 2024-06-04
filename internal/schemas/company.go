package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Name  string
	TaxID string
	Email string
	Phone string
}

type CompanyResponse struct {
	ID        uint           `json:"id"`
	Name      string         `json:"name"`
	TaxID     string         `json:"taxId"`
	Email     string         `json:"email"`
	Phone     string         `json:"phone"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty"`
}
