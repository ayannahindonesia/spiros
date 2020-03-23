package helper

import "github.com/labstack/echo"

// echoResp type
type echoResp struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}

// ReturnJSONresp returns response with format
func ReturnJSONresp(c echo.Context, httpcode int, code string, message string, details interface{}) error {
	if len(code) <= 0 {
		code = "0000"
	}
	x := echoResp{
		Code:    code,
		Message: message,
		Details: details,
	}

	return c.JSON(httpcode, x)
}
