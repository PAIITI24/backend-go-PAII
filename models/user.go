package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Id        int `gorm:"primaryKey,autoIncrement"`
	Name      string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
