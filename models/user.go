package models

import "gorm.io/gorm"


type User struct {
	gorm.Model
	Nama          string `gorm:"size:255"`
	NIK           string `gorm:"size:16"`
	NomorHP       string `gorm:"size:16"`
	NomorRekening string `gorm:"size:255"`
}
