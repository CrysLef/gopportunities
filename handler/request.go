package handler

import (
	"fmt"
	"reflect"
)

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func errParamIsRequired(param string, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", param, typ)
}

func typeOf(v interface{}) string {
	return fmt.Sprint(reflect.TypeOf(v))
}

func (req *CreateOpeningRequest) Validate() error {

	if req.Role == "" && req.Company == "" && req.Link == "" && req.Location == "" && req.Salary <= 0 && req.Remote == nil {
		return fmt.Errorf("request body empty or malformed")
	}
	if req.Role == "" {
		errParamIsRequired("role", typeOf(req.Role))
	}
	if req.Company == "" {
		errParamIsRequired("company", typeOf(req.Company))
	}
	if req.Location == "" {
		errParamIsRequired("location", typeOf(req.Location))
	}
	if req.Link == "" {
		errParamIsRequired("link", typeOf(req.Link))
	}
	if req.Remote == nil {
		errParamIsRequired("remote", typeOf(req.Remote))
	}
	if req.Salary <= 0 {
		errParamIsRequired("salary", typeOf(req.Salary))
	}
	return nil
}