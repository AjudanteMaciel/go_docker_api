package schemas

import (
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Name      string     `gorm:"not null"`
	Employees []Employee `gorm:"foreignKey:EmployeeID"`
}
