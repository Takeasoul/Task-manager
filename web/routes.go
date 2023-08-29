package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func TaskPage(c echo.Context) error {
	return c.Render(http.StatusOK, "tasks.html", nil)
}
