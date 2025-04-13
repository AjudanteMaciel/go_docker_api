package schemas

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name    string `gorm:"not null"`
	Email   string `gorm:"not null;unique"`
	RoleID  *uint
	StateID *uint
}

type EmployeeResponse struct {
	gorm.Model
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	RoleID  *uint  `json:"role_id"`
	StateID *uint  `json:"state_id"`
}
