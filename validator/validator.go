package validator

import (
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/cast"
	"go.uber.org/multierr"
	sortModels "phantom/dataLayer"
	"phantom/dataLayer/cacheDaos"
	"reflect"
	"strings"
)

var ErrParamNotPtr = errors.New("parameter must be a pointer")
var ErrPtrNotStruct = errors.New("pointer must be of a struct")
var tagName = "validate"

// Daos used in validation
var categoryCacheDao *cacheDaos.AllCategoryIdsRedisDao
var propertyValueIdCacheDao *cacheDaos.PropertyValueIdToPropertyIdRedisDao
var productCacheDao *cacheDaos.AllProductIdsRedisDao

func Init(pool *redis.Pool) {
	categoryCacheDao = &cacheDaos.AllCategoryIdsRedisDao{Pool: pool}
	propertyValueIdCacheDao = &cacheDaos.PropertyValueIdToPropertyIdRedisDao{Pool: pool}
	productCacheDao = &cacheDaos.AllProductIdsRedisDao{Pool: pool}
}

func Validate(data interface{}) error {
	// Check that data kind is a pointer
	dataType := reflect.TypeOf(data)
	if dataType.Kind() != reflect.Ptr {
		return ErrParamNotPtr
	}
	dataType = dataType.Elem()
	dataValue := reflect.ValueOf(data).Elem()

	// Check that data points to a struct
	if dataValue.Kind() != reflect.Struct {
		return ErrPtrNotStruct
	}

	// For each field we validate it and append the field errors to structErrArr
	structErrArr := make([]error, 0)
	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)
		fieldValue := dataValue.Field(i)
		structFieldData := makeStructFieldData(field, fieldValue)
		fieldError := validateField(structFieldData)
		if fieldError != nil {
			structErrArr = append(structErrArr, *fieldError)
		}
	}
	combinedStructError := multierr.Combine(structErrArr...)
	return combinedStructError
}

func validateField(structFieldData StructFieldData) *error {
	// Find conditions for field
	validateStr := structFieldData.Tag.Get(tagName)
	validateConditions := strings.Split(validateStr, ",")

	// Iterate over conditions and collect their errors in fieldErrArr
	fieldErrArr := make([]error, 0)
	for _, condition := range validateConditions {
		conditionErr := validateCondition(structFieldData, condition)
		if conditionErr != nil {
			fieldErrArr = append(fieldErrArr, *conditionErr)
		}
	}
	if len(fieldErrArr) > 0 {
		combinedFieldErr := multierr.Combine(fieldErrArr...)
		return &combinedFieldErr
	}
	return nil
}

func validateCondition(fieldData StructFieldData, condition string) *error {
	var conditionErr *error
	switch {
	case condition == "required":
		conditionErr = required(fieldData)
	case condition == "category_id":
		conditionErr = isValidCategoryId(fieldData)
	case condition == "sort_id":
		conditionErr = isValidSortId(fieldData)
	case condition == "property_ids":
		conditionErr = isValidPropertyIds(fieldData)
	case condition == "product_id":
		conditionErr = isValidProductId(fieldData)
	}
	return conditionErr
}

/* Condition functions
https://github.com/go-playground/validator/blob/d07eb88fb04047e450e834dd461210095fc28d6a/baked_in.go
*/

func required(fieldData StructFieldData) *error {
	if fieldData.Value.IsZero() {
		err := errors.New(fmt.Sprintf("%s must be present", fieldData.Name))
		return &err
	}
	return nil
}

func isValidCategoryId(fieldData StructFieldData) *error {
	categoryId, castErr := cast.ToInt64E(fieldData.Value.Interface())
	if castErr != nil {
		fieldTypeErr := errors.New(fmt.Sprintf("%s must be an integer", fieldData.Name))
		return &fieldTypeErr
	}
	isMember, cacheErr := categoryCacheDao.IsMember(categoryId)
	if cacheErr != nil {
		err := errors.New(fmt.Sprintf("%s isn't valid. Something went wrong", fieldData.Name))
		return &err
	}

	if !isMember {
		err := errors.New(fmt.Sprintf("%s must be a valid category id", fieldData.Name))
		return &err
	}
	return nil
}

func isValidSortId(fieldData StructFieldData) *error {
	sortId, castErr := cast.ToInt64E(fieldData.Value.Interface())
	if castErr != nil {
		fieldTypeErr := errors.New(fmt.Sprintf("%s must be an integer", fieldData.Name))
		return &fieldTypeErr
	}
	isSortIdValid := false
	for _, sortMethod := range sortModels.AllSortMethods {
		if sortMethod.Id == sortId {
			isSortIdValid = true
		}
	}

	if !isSortIdValid {
		err := errors.New(fmt.Sprintf("%s must be a valid sort id", fieldData.Name))
		return &err
	}
	return nil
}

func isValidPropertyIds(fieldData StructFieldData) *error {
	propertyIdsArr, ok := fieldData.Value.Interface().([]int64)
	if !ok {
		fieldTypeErr := errors.New(fmt.Sprintf("%s must be an integer array", fieldData.Name))
		return &fieldTypeErr
	}
	propertyIdErrArr := make([]error, 0)
	for _, propertyId := range propertyIdsArr {
		valid, cacheErr := propertyValueIdCacheDao.IsPropertyIdValid(propertyId)
		if cacheErr != nil {
			err := errors.New(fmt.Sprintf("%s isn't valid. Something went wrong", fieldData.Name))
			return &err
		}
		if !valid {
			notValidErr := errors.New(fmt.Sprintf("%d must be a valid property id", propertyId))
			propertyIdErrArr = append(propertyIdErrArr, notValidErr)
		}
	}

	combinedErr := multierr.Combine(propertyIdErrArr...)
	if combinedErr != nil {
		return &combinedErr
	}
	return nil
}

func isValidProductId(fieldData StructFieldData) *error {
	productId, castErr := cast.ToInt64E(fieldData.Value.Interface())
	if castErr != nil {
		fieldTypeErr := errors.New(fmt.Sprintf("%s must be an integer", fieldData.Name))
		return &fieldTypeErr
	}
	isMember, cacheErr := productCacheDao.IsValidProductId(productId)
	if cacheErr != nil {
		err := errors.New(fmt.Sprintf("%s isn't valid. Something went wrong", fieldData.Name))
		return &err
	}

	if !isMember {
		err := errors.New(fmt.Sprintf("%s must be a valid product id", fieldData.Name))
		return &err
	}

	return nil
}
