package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCompanyHandler(ctx *gin.Context) {
	request := CreateCompanyRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := db.Create(&request).Error; err != nil {
		logger.Errorf("error creating company: %v", err.Error())
		return
	}
}
