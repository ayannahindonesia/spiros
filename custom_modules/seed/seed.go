package seed

import (
	"spiros/db"
	"spiros/models"

	"golang.org/x/crypto/bcrypt"
)

// Seed inserts dummy datas
func Seed() {
	clientsecret, _ := bcrypt.GenerateFromPassword([]byte("clientsecret"), bcrypt.DefaultCost)
	client := models.Client{
		Key:    "clientkey",
		Secret: string(clientsecret),
	}
	db.DB.Create(&client)

	adminpass, _ := bcrypt.GenerateFromPassword([]byte("rahasia123"), bcrypt.DefaultCost)
	admin := models.User{
		Username: "admin",
		Password: string(adminpass),
	}
	db.DB.Create(&admin)
}

// Unseed removes dummy datas
func Unseed() {

}
