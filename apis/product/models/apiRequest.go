package models

import (
	"errors"
	"github.com/gin-gonic/gin"
	"phantom/apis/apiCommons"
	"phantom/validator"
)

type ApiRequest struct {
	ProductId int64 `form:"product_id" validate:"required,product_id"`
}

func ReadApiRequestModel(ctx *gin.Context) (*ApiRequest, error) {
	var apiRequest ApiRequest

	// Bind the query to our request model
	queryParamsReadErr := ctx.BindQuery(&apiRequest)

	// Binding error handling
	if queryParamsReadErr != nil {
		msg := "something went wrong reading the request model"
		logData := apiCommons.NewApiErrorLogData(ctx, msg, queryParamsReadErr)
		apiCommons.LogApiError(logData)
		return nil, errors.New(msg)
	}

	// Check validation of the request model
	validationErr := validator.Validate(&apiRequest)
	if validationErr != nil {
		msg := "request model isn't valid"
		logData := apiCommons.NewApiErrorLogData(ctx, msg, validationErr)
		apiCommons.LogApiError(logData)
		return nil, validationErr
	}

	return &apiRequest, nil
}
