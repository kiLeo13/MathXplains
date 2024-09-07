package service

import (
	"fmt"
)

var (
	ErrorMalformedJSON  = NewError(400, "Malformed JSON body")
	ErrorInternalServer = NewError(500, "Internal server error")
	ErrorMissingAdmin   = NewError(403, "Missing admin privileges")

	ErrorUserExists           = NewError(400, "User already exists")
	ErrorUserNotFound         = NewError(404, "User not found")
	ErrorInvalidNameRange     = NewError(400, "Name is out of range: (%d - %d)", nameMinLength, nameMaxLength)
	ErrorInvalidEmailRange    = NewError(400, "Email is out of range: (%d - %d)", emailMinLength, emailMaxLength)
	ErrorPasswordCase         = NewError(400, "Passwords must have least one uppercase and lowercase character")
	ErrorPasswordSpecialChar  = NewError(400, "Passwords must have at least one special character")
	ErrorPasswordNumbers      = NewError(400, "Passwords must have at least one number")
	ErrorInvalidPasswordRange = NewError(400, "Password is out of range: (%d - %d)", passwordMinLength, passwordMaxLength)

	ErrorTooManyAppointments       = NewError(400, "Too many active appointments")
	ErrorAppointmentTooOldToDelete = NewError(400, "This appointment is too old to be deleted (max: %d seconds in the past).", int32(maxDeletionPeriod.Seconds()))
	ErrorInvalidTopicRange         = NewError(400, "Topic is out of range: (%d - %d)", minTopicLength, maxTopicLength)
	ErrorInvalidDescriptionRange   = NewError(400, "Description is out of range: (%d - %d)", minDescLength, maxDescLength)
	ErrorAppointmentNotFound       = NewError(404, "Appointment not found")

	ErrorSubjectDoesNotExist = NewError(400, "Subject does not exist")
	ErrorSubjectUnavailable  = NewError(400, "Subject is unavailable")

	ErrorIncorrectDateFormat = NewError(400, "Incorrect date format, expected: 'yyyy-mm-dd'")
)

type APIError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func NewError(status int, msg string, args ...any) *APIError {
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	return &APIError{Status: status, Message: msg}
}

func ErrorInvalidPattern(paramName string) *APIError {
	return NewError(400, `Parameter "%s" is not in the expected pattern`, paramName)
}

func ErrorInvalidPathParam(path, dataType string) *APIError {
	return NewError(400, `Invalid path parameter: "%s" should be of type %s`, path, dataType)
}

func ErrorParamNotProvided(name string) *APIError {
	return NewError(400, `Parameter "%s" was not provided`, name)
}
