package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Name     string `json:"name" form:"name"`
	Telp     string `json:"telp" form:"telp"`
	Alamat   string `json:"alamat" form:"alamat"`
	Email    string `json:"email" form:"email"`
}
