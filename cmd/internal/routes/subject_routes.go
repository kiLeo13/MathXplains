package routes

import (
	"MathXplains/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetSubjects(c echo.Context) error {
	available := c.QueryParam("available") == "true"
	subjects := service.GetSubjects(available)
	resp := R{
		"subjects": subjects,
	}

	return c.JSON(http.StatusOK, &resp)
}
