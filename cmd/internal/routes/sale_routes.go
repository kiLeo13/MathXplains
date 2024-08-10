package routes

import (
	"MathXplains/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetSales(c echo.Context) error {
	count := service.GetSalesCount()
	return c.JSON(http.StatusOK, &R{service.SalesKeyName: count})
}

func PatchSalesCount(c echo.Context) error {
	body := make(map[string]int)
	userId := c.Request().Header.Get("Sub")
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, service.ErrorMalformedJSON)
	}

	value, ok := body["amount"]
	admin := service.IsAdmin(userId)
	if !admin {
		return c.JSON(http.StatusForbidden, service.ErrorMissingAdmin)
	}

	if !ok {
		return c.JSON(http.StatusBadRequest, service.ErrorParamNotProvided("amount"))
	}

	if value == 0 {
		return c.JSON(http.StatusOK, &R{service.SalesKeyName: service.GetSalesCount()})
	}

	newCount, err := service.UpdateSalesCount(value)
	if err != nil {
		return c.JSON(err.Code, err)
	}
	return c.JSON(http.StatusOK, &R{service.SalesKeyName: newCount})
}
