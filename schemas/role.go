package schemas

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Admin     bool       `gorm:"not null"`
	Name      string     `gorm:"not null;unique"`
	Employees []Employee `gorm:"foreignKey:RoleID"`
}

type RoleResponse struct {
	gorm.Model
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
}
