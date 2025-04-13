package employee

import "go_docker_api/schemas"

type CreateEmployeeResponse struct {
	Message string                   `json:"message"`
	Data    schemas.EmployeeResponse `json:"data"`
}

type DeleteEmployeeResponse struct {
	Message string                   `json:"message"`
	Data    schemas.EmployeeResponse `json:"data"`
}
type ShowEmployeeResponse struct {
	Message string                   `json:"message"`
	Data    schemas.EmployeeResponse `json:"data"`
}
type ListOpeningsResponse struct {
	Message string                     `json:"message"`
	Data    []schemas.EmployeeResponse `json:"data"`
}
type UpdateEmployeeResponse struct {
	Message string                   `json:"message"`
	Data    schemas.EmployeeResponse `json:"data"`
}
