package routes

import (
	"MathXplains/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetDollarExchange(c echo.Context) error {
	value, err := service.GetCurrentDollarExchange()
	if err != nil {
		return c.JSON(err.Code, err)
	}
	resp := R{service.DollarExchangeRateKey: value}

	return c.JSON(http.StatusOK, &resp)
}
