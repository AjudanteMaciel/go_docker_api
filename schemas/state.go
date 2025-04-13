package schemas

import (
	"gorm.io/gorm"
)

type State struct {
	gorm.Model
	Name         string     `gorm:"not null;unique"`
	Abbreviation string     `gorm:"not null"`
	Employee     []Employee `gorm:"foreignKey:StateID"`
}
