package handler

import (
	"fmt"
)

func errParameterIsRequired(name, typ string) error {
	return fmt.Errorf("%s parameter is required (type %s)", name, typ)
}

type Opening struct {
	Role     string  `json:"role"`
	Company  string  `json:"company"`
	Location string  `json:"location"`
	Remote   *bool   `json:"remote"`
	Link     string  `json:"link"`
	Salary   float64 `json:"salary"`
}

func (r *Opening) Validate() error {
	if r.Role == "" && r.Location == "" && r.Company == "" && r.Link == "" && r.Salary <= 0 && r.Remote == nil {
		return fmt.Errorf("malformed or empty request body for Opening")
	}
	if len(r.Role) == 0 {
		return errParameterIsRequired("role", "string")
	}
	if len(r.Company) == 0 {
		return errParameterIsRequired("company", "string")
	}
	if len(r.Location) == 0 {
		return errParameterIsRequired("company", "string")
	}
	if r.Remote == nil {
		return errParameterIsRequired("remote", "boolean")
	}
	if r.Salary <= 0 {
		return errParameterIsRequired("salary", "float")
	}
	if len(r.Link) == 0 {
		return errParameterIsRequired("link", "string")
	}
	return nil
}
