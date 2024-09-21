package routes

import (
	"MathXplains/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetNotes(c echo.Context) error {
	profile := c.QueryParam("profile")
	if len(profile) == 0 {
		return c.JSON(http.StatusBadRequest, service.ErrorProfileNotProvided)
	}

	notes, err := service.GetNotes(profile)
	if err != nil {
		return c.JSON(err.Status, err)
	}

	resp := R{
		"notes": notes,
	}
	return c.JSON(http.StatusOK, &resp)
}

func CreateNote(c echo.Context) error {
	var note service.NoteCreateDTO
	if err := c.Bind(&note); err != nil {
		return c.JSON(http.StatusBadRequest, service.ErrorMalformedJSON)
	}

	newNote, err := service.CreateNote(&note)
	if err != nil {
		return c.JSON(err.Status, err)
	}
	return c.JSON(http.StatusOK, newNote)
}

func PutNote(c echo.Context) error {
	idIn := c.Param("id")
	var note service.NoteCreateDTO
	if err := c.Bind(&note); err != nil {
		return c.JSON(http.StatusBadRequest, service.ErrorMalformedJSON)
	}

	id, err := ToInt("id", idIn)
	if err != nil {
		return c.JSON(err.Status, err)
	}

	newNote, err := service.PutNote(id, &note)
	if err != nil {
		return c.JSON(err.Status, err)
	}
	return c.JSON(http.StatusOK, newNote)
}

func DeleteNote(c echo.Context) error {
	idIn := c.Param("id")
	id, err := ToInt("id", idIn)
	if err != nil {
		return c.JSON(err.Status, err)
	}

	err = service.DeleteNote(id)
	if err != nil {
		return c.JSON(err.Status, err)
	}

	return c.NoContent(http.StatusOK)
}
