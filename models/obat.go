package models

import (
	"gorm.io/gorm"
	"time"
)

type Obat struct {
	gorm.Model
	Id            int `gorm:"primaryKey,autoIncrement"`
	NamaObat      string
	DosisObat     string
	BentukSediaan string
	Harga         float32
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}
