package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/storify/backend/internal/schemas"
)

func GetCompanyHandler(ctx *gin.Context) {
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

	sendSuccess(ctx, "company retrieved", company)
}
