package employee

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateEmployeeRequest struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required,email"`
	RoleID  *uint  `json:"role_id"`
	StateID *uint  `json:"state_id"`
}

func (request *CreateEmployeeRequest) Validate() error {
	if request.Name == "" {
		return errParamIsRequired("role", "string")
	}
	if request.Email == "" {
		return errParamIsRequired("email", "string")
	}

	return nil
}

type UpdateEmployeeRequest struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required,email"`
	RoleID  *uint  `json:"role_id"`
	StateID *uint  `json:"state_id"`
}

func (request *UpdateEmployeeRequest) Validate() error {
	if request.Name == "" || request.Email == "" {
		return fmt.Errorf("at least one of the fields is required")
	}

	return nil
}
