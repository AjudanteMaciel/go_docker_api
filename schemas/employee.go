package schemas

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name    string `gorm:"not null"`
	Email   string `gorm:"not null;unique"`
	RoleID  uint   `gorm:"not null"`
	StateID uint   `gorm:"not null"`
}
