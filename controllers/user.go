package controllers

import (
	"PAII_Back-End/config"
	"PAII_Back-End/models"
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(ctx *fiber.Ctx) error {

	// recieve body
	fInput := struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}{}

	dec := json.NewDecoder(bytes.NewReader(ctx.Body()))
	err := dec.Decode(&fInput)
	if err != nil {
		return err
	}

	ctx.Write([]byte("hello, your name is " + fInput.Name))
	db := config.DB()

	dInput := models.User{
		Name:     fInput.Name,
		Email:    fInput.Email,
		Password: fInput.Password,
		Role:     fInput.Role,
	}

	res := db.Create(&dInput)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
