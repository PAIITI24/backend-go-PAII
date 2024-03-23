package models

type ResObatKategoriObat struct {
	Id             int `gorm:"primaryKey"`
	IdObat         int
	IdKategoriObat int
	Obat           Obat         `gorm:"foreignKey:IdObat;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	KategoriObat   KategoriObat `gorm:"foreignKey:IdKategoriObat;references:Id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}
