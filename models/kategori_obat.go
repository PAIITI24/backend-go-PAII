package models

import (
	"gorm.io/gorm"
	"time"
)

type KategoriObat struct {
	gorm.Model
	Id               int `gorm:"primaryKey,autoIncrement"`
	NamaKategoriObat string
	CreatedAt        time.Time `gorm:"autoCreateTime"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime"`
}
