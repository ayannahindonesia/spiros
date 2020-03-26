package handlers

import (
	"fmt"
	"net/http"
	"spiros/db"
	"spiros/models"
	"spiros/server/helper"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// ViewCurrentUser shows current user's datas
func ViewCurrentUser(c echo.Context) error {
	defer c.Request().Body.Close()

	type ViewUser struct {
		Username string `json:"username"`
	}

	userToken := c.Get("user").(*jwt.Token)
	tokenClaims := userToken.Claims.(jwt.MapClaims)
	userID, _ := strconv.ParseUint(tokenClaims["jti"].(string), 10, 64)

	user := models.User{}
	err := db.DB.First(&user, "id = ?", userID).Error
	if err != nil {
		return helper.ReturnJSONresp(c, http.StatusUnauthorized, "0002", "Error finding user. please re-login", map[string]interface{}{
			"error": fmt.Sprint(err),
		})
	}

	return helper.ReturnJSONresp(c, http.StatusOK, "0000", "Success", ViewUser{Username: user.Username})
}
