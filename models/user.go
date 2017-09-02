package models

import (
	"github.com/jinzhu/gorm"
)

// User estructura de la tabla usuario
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"not null; type:varchar(25)"`
	Password string `json:"password" gorm:"not null; type:varchar(256)"`
	Level    string `json:"level" gorm:"not null; type:varchar(20); default:'ASISTENTE'"`
}

// Users slice de usuarios
type Users []User
