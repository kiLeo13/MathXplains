package routes

import (
	"MathXplains/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetAppointments(c echo.Context) error {
	all := c.QueryParam("all") == "true"
	status := c.QueryParam("status")
	userId := c.Request().Header.Get("Sub")

	if all && !service.IsAdmin(userId) {
		return c.JSON(http.StatusForbidden, service.ErrorMissingAdmin)
	}

	apptms, err := service.GetAppointments(all, status, userId)
	if err != nil {
		return c.JSON(err.Code, err)
	}

	resp := R{
		"appointments": apptms,
		"active_count": countActiveAppointments(apptms),
		"max":          service.MaxActiveAppointments,
	}
	return c.JSON(http.StatusOK, &resp)
}

func CreateAppointment(c echo.Context) error {
	userId := c.Request().Header.Get("Sub")
	body := service.AppointmentCreateDTO{}
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, service.ErrorMalformedJSON)
	}

	appt, err := service.CreateAppointment(userId, &body)
	if err != nil {
		return c.JSON(err.Code, err)
	}
	return c.JSON(http.StatusOK, appt)
}

func DeleteAppointment(c echo.Context) error {
	id := c.Param("id")
	userId := c.Request().Header.Get("Sub")

	err := service.DeleteAppointment(userId, id)
	if err != nil {
		return c.JSON(err.Code, err)
	}
	return c.NoContent(http.StatusNoContent)
}

func countActiveAppointments(apptms []*service.AppointmentDTO) int {
	count := 0
	for _, a := range apptms {
		if a.IsActive {
			count++
		}
	}
	return count
}
