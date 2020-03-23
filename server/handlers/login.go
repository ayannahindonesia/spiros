package handlers

import (
	"fmt"
	"net/http"
	"os"
	"spiros/db"
	"spiros/models"
	"spiros/server/helper"
	"strconv"

	"github.com/labstack/echo"
)

// Login user using username and password
func Login(c echo.Context) error {
	defer c.Request().Body.Close()

	type Body struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
	body := Body{}

	err := c.Bind(&body)
	if err != nil {
		panic(err)
	}

	user := models.User{}
	db.DB.Model(&models.User{}).Find(&user, models.User{Username: body.Username})

	expiresIn, err := strconv.ParseInt(os.Getenv("SPIROS_JWT_EXPIRES"), 10, 64)
	if err != nil {
		return helper.ReturnJSONresp(c, http.StatusInternalServerError, "0000", "Error", map[string]interface{}{
			"error": err,
		})
	}

	token, err := helper.GenerateJWTtoken(fmt.Sprint(user.ID))
	if err != nil {
		return helper.ReturnJSONresp(c, http.StatusInternalServerError, "0000", "Error", map[string]interface{}{
			"error": err,
		})
	}

	return helper.ReturnJSONresp(c, http.StatusOK, "0000", "Success", map[string]interface{}{
		"expires_in": expiresIn,
		"token":      token,
	})
}
