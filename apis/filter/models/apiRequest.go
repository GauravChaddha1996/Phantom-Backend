package models

import (
	"errors"
	"github.com/gin-gonic/gin"
	"phantom/apis/apiCommons"
	"phantom/validator"
)

type ApiRequest struct {
	CategoryId          int64   `form:"category_id"`
	PropertyValueIds    []int64 `form:"property_value_ids"`
	SortId              int64   `form:"sort_id"`
	PropertyValueIdsMap map[int64]bool
}

func ReadApiRequestModel(ctx *gin.Context) (*ApiRequest, error) {
	var apiRequest ApiRequest
	queryParamsReadErr := ctx.BindQuery(&apiRequest)
	if queryParamsReadErr != nil {
		msg := "something went wrong reading the request model"
		logData := apiCommons.NewApiErrorLogData(ctx, msg, queryParamsReadErr)
		apiCommons.LogApiError(logData)
		return nil, errors.New(msg)
	}

	validationErr := validator.Validate(&apiRequest)
	if validationErr != nil {
		msg := "request model isn't valid"
		logData := apiCommons.NewApiErrorLogData(ctx, msg, validationErr)
		apiCommons.LogApiError(logData)
		return nil, validationErr
	}

	// Fill some useful map for later
	apiRequest.PropertyValueIdsMap = make(map[int64]bool, 0)
	for _, propertyValueId := range apiRequest.PropertyValueIds {
		apiRequest.PropertyValueIdsMap[propertyValueId] = true
	}
	return &apiRequest, nil
}
