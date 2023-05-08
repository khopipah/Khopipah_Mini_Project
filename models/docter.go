package models

import (
	"gorm.io/gorm"
)

type Docter struct {
	gorm.Model
	Name      string `json:"name" form:"name"`
	Spesialis string `json:"spesialis" form:"spesialis"`
	Gender    string `json:"gender" form:"gender"`
	Telp      string `json:"telp" form:"telp"`
	Alamat    string `json:"alamat" form:"alamat"`
	Pasien    []Pasien
}
