package models

import (
	"gorm.io/gorm"
	"time"
)

type KategoriProduk struct {
	gorm.Model
	Id           int `gorm:"PrimaryKey"`
	NamaKategori string
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}
