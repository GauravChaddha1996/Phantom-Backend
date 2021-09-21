package validator

import (
	"errors"
	"fmt"
	"go.uber.org/multierr"
	"reflect"
	"strings"
)

var ErrParamNotPtr = errors.New("parameter must be a pointer")
var ErrPtrNotStruct = errors.New("pointer must be of a struct")
var tagName = "validate"

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
