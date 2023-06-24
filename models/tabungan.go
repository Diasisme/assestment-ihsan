package models

import "gorm.io/gorm"

type Tabungan struct {
	gorm.Model
	NomorRekening string `gorm:"size:255"`
	Saldo         float64
	KodeTransaksi string `gorm:"size:2"`
}
