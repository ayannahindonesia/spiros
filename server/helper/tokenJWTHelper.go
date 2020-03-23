package helper

import (
	"log"
	"os"
	"spiros/db"
	"spiros/models"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

// GenerateJWTtoken generates jwt token
func GenerateJWTtoken(id string) (string, error) {
	expiresIn, err := strconv.Atoi(os.Getenv("SPIROS_JWT_EXPIRES"))
	if err != nil {
		return "", err
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Duration(expiresIn) * time.Second).Unix(),
		Id:        id,
	})

	token, err := rawToken.SignedString([]byte(os.Getenv("SPIROS_JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return token, nil
}

// ValidateClient check client from client table in database
func ValidateClient(key, secret string, c echo.Context) (bool, error) {
	client := models.Client{Key: key}
	err := db.DB.First(&client, client).Error
	if err != nil {
		log.Println(err)
		return false, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(client.Secret), []byte(secret))
	if err != nil {
		log.Println(err)
		return false, err
	}
	return true, nil
}
