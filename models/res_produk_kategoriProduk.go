package models

import "gorm.io/gorm"

type ResProdukKategoriProduk struct {
	gorm.Model
	IdProduk         int
	IdKategoriProduk int
	Produk           Produk         `gorm:"foreignKey:IdProduk;references:Id;constraint:onUpdate:CASCADE,OnDelete:SET NULL"`
	KategoriProduk   KategoriProduk `gorm:"foreignKey:IdKategoriProduk;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
