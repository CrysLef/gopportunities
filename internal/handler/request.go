package handler

import (
	"fmt"
	"reflect"
)

type HandlerOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   bool   `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func errParamIsRequired(param string, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", param, typ)
}

func InfoStruct(s interface{}) (reflect.StructField, int, reflect.Value) {
	v := reflect.ValueOf(s).Elem()
	numField := v.NumField()
	var field reflect.StructField

	return field, numField, v
}

func (req *HandlerOpeningRequest) ValidateCreate() error {

	field, _, v := InfoStruct(req)

	for i := 0; i < v.Len(); i++ {
		field = v.Type().Field(i)
		fieldName := field.Name
		fieldType := field.Type.Name()
		fieldValue := v.Field(i)

		if fieldType == "string" {
			if fieldValue.String() == "" {
				return errParamIsRequired(fieldName, fieldType)
			}
		}

		if fieldType == "int64" {
			if fieldValue.IsZero() {
				return errParamIsRequired(fieldName, fieldType)
			}
		}
	}
	return nil
}

func (req *HandlerOpeningRequest) ValidateUpdate() error {

	field, numField, v := InfoStruct(req)
	for i := 0; i < numField; i++ {
		field = v.Type().Field(i)
		fieldValue := v.Field(i)
		fieldType := field.Type.Name()

		if fieldType == "string" {
			if fieldValue.String() != "" {
				return nil
			}
		} else if fieldType == "int64" {
			if fieldValue.Int() > 0 {
				return nil
			}
		} else if fieldType == "bool" {
			if fieldValue.Bool() {
				return nil
			}
		}
	}

	return fmt.Errorf("at least one field must be provided")
}
