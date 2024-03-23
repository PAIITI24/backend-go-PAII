package main

import (
	"PAII_Back-End/config"
	"PAII_Back-End/models"
	"fmt"
)

func main() {
	db := config.DB()

	err := db.AutoMigrate(
		&models.User{},
		&models.Obat{},
		&models.KategoriObat{},
		&models.ResObatKategoriObat{},
	)

	if err != nil {
		fmt.Println("Migration error: ", err)
		return
	}

	fmt.Println("Migration successful")
}
