package router

import (
	"os"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// NewRouter func
func NewRouter() *echo.Echo {
	e := echo.New()

	// url rewrite
	url := strings.Split(os.Getenv("SPIROS_URL_REWRITE"), ",")
	x := make(map[string]string)
	for _, v := range url {
		if len(v) > 0 {
			x[v] = "/$1"
		}
	}
	if len(x) > 0 {
		e.Pre(middleware.Rewrite(x))
	}

	ClientGroup(e)

	return e
}
