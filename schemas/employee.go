package schemas

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name  string `gorm:"not null"`
	Roles []Role `gorm:"many2many:employee_roles;"`
}
