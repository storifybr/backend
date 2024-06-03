package handler

import (
	"github.com/storify/backend/configs"
	"gorm.io/gorm"
)

var (
	logger *configs.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = configs.GetLogger("handler")
	db = configs.GetDatabase()
}
