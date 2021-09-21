package models

import (
	"errors"
	"github.com/gin-gonic/gin"
	"phantom/apis/apiCommons"
	"phantom/validator"
)

type ApiRequestModel struct {
	CategoryId       int64   `form:"category_id"`
	PropertyValueIds []int64 `form:"property_value_ids"`
}

func ReadApiRequestModel(ctx *gin.Context) (*ApiRequestModel, error) {
	var apiRequestModel ApiRequestModel
	queryParamsReadErr := ctx.BindQuery(&apiRequestModel)
	if queryParamsReadErr != nil {
		msg := "something went wrong reading the request model"
		apiErrorLogData := apiCommons.NewApiErrorLogData(ctx, msg, queryParamsReadErr)
		apiCommons.LogApiError(apiErrorLogData)
		return nil, errors.New(msg)
	}

	validationErr := validator.Validate(&apiRequestModel)
	if validationErr != nil {
		msg := "request model isn't valid"
		apiErrorLogData := apiCommons.NewApiErrorLogData(ctx, msg, validationErr)
		apiCommons.LogApiError(apiErrorLogData)
		return nil, validationErr
	}

	return &apiRequestModel, nil
}
