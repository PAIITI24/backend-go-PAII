package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strconv"
)

var (
	konfig_db = struct {
		username string
		password string
		hostname string
		port     int
		dbname   string
	}{
		"root", "", "localhost", 3306, "db_golang_pa",
	}
)

func DB() *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: konfig_db.username + ":" + konfig_db.password + "@tcp(" + konfig_db.hostname + ":" + strconv.Itoa(konfig_db.port) + ")/" + konfig_db.dbname,
	}))

	if err != nil {
		log.Fatal(err)
	}

	return db
}
