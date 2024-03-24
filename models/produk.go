package models

import (
	"gorm.io/gorm"
	"time"
)

type Produk struct {
	gorm.Model
	Id             int `gorm:"primaryKey"`
	NamaProduk     string
	Harga          float64
	KategoriProduk string
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}
