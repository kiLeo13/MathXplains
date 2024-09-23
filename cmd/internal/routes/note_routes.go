package routes

import (
	"MathXplains/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetNote(c echo.Context) error {
	profile := c.QueryParam("profile")
	idIn := c.Param("id")
	id, err := ToInt("id", idIn)
	if err != nil {
		return c.JSON(err.Status, err)
	}

	note, err := service.GetNote(profile, id)
	if err != nil {
		return c.JSON(err.Status, err)
	}
	return c.JSON(http.StatusOK, note)
}

func GetNotesSummary(c echo.Context) error {
	profile := c.QueryParam("profile")
	if len(profile) == 0 {
		return c.JSON(http.StatusBadRequest, service.ErrorProfileNotProvided)
	}

	notes, err := service.GetNotesSummary(profile)
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

func UpdateNote(c echo.Context) error {
	idIn := c.Param("id")
	profile := c.Request().Header.Get("Profile")
	var note service.NoteCreateDTO
	if err := c.Bind(&note); err != nil {
		return c.JSON(http.StatusBadRequest, service.ErrorMalformedJSON)
	}

	id, err := ToInt("id", idIn)
	if err != nil {
		return c.JSON(err.Status, err)
	}

	newNote, err := service.PutNote(profile, id, &note)
	if err != nil {
		return c.JSON(err.Status, err)
	}
	return c.JSON(http.StatusOK, newNote)
}

func DeleteNote(c echo.Context) error {
	idIn := c.Param("id")
	profile := c.Request().Header.Get("Profile")
	id, err := ToInt("id", idIn)
	if err != nil {
		return c.JSON(err.Status, err)
	}

	err = service.DeleteNote(profile, id)
	if err != nil {
		return c.JSON(err.Status, err)
	}
	return c.NoContent(http.StatusOK)
}
