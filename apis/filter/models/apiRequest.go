package models

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"phantom/apis/apiCommons"
	"phantom/validator"
)

type ApiRequest struct {
	CategoryId          int64  `form:"category_id"`
	PropertyValueIds    string `form:"property_value_ids"`
	SortId              int64  `form:"sort_id"`
	PropertyValueIdsMap map[int64]bool
}

func ReadApiRequestModel(ctx *gin.Context) (*ApiRequest, error) {
	var apiRequest ApiRequest

	// Bind the query to our request model
	queryParamsReadErr := ctx.BindQuery(&apiRequest)
	// Some sad manual binding
	var propertyValueIdArr []int64
	var propertyValueIdUnmarshalErr error
	if len(apiRequest.PropertyValueIds) > 0 {
		propertyValueIdUnmarshalErr = json.Unmarshal([]byte(apiRequest.PropertyValueIds), &propertyValueIdArr)
	}

	// Binding error handling
	var finalBindingErr error = nil
	if queryParamsReadErr != nil {
		finalBindingErr = queryParamsReadErr
	} else if propertyValueIdUnmarshalErr != nil {
		finalBindingErr = propertyValueIdUnmarshalErr
	}
	if finalBindingErr != nil {
		msg := "something went wrong reading the request model"
		logData := apiCommons.NewApiErrorLogData(ctx, msg, finalBindingErr)
		apiCommons.LogApiError(logData)
		return nil, errors.New(msg)
	}

	// Fill some useful map for later
	apiRequest.PropertyValueIdsMap = make(map[int64]bool, 0)
	for _, propertyValueId := range propertyValueIdArr {
		apiRequest.PropertyValueIdsMap[propertyValueId] = true
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
