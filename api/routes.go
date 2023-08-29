package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetTasks(c echo.Context) error {
	return c.JSON(http.StatusOK, Tasks)
}
