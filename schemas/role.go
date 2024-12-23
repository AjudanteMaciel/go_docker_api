package schemas

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Admin     bool       `gorm:"not null"`
	Name      string     `gorm:"not null;unique"`
	Employees []Employee `gorm:"many2many:employee_roles;"`
}
