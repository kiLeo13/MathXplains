package service

import (
	"fmt"
)

var (
	ErrorMalformedJSON  = NewError(400, "Malformed JSON body")
	ErrorInternalServer = NewError(500, "Internal server error")
	ErrorUserExists     = NewError(400, "User already exists")
	ErrorUserNotFound   = NewError(404, "User not found")
	ErrorMissingAdmin   = NewError(403, "Missing admin privileges")

	ErrorTooManyAppointments     = NewError(400, "Too many active appointments")
	ErrorInvalidTopicRange       = NewError(400, "Topic is out of range: (%d - %d)", minTopicLength, maxTopicLength)
	ErrorInvalidDescriptionRange = NewError(400, "Description is out of range: (%d - %d)", minDescLength, maxDescLength)
	ErrorSubjectDoesNotExist     = NewError(400, "Subject does not exist")
	ErrorSubjectUnavailable      = NewError(400, "Subject is unavailable")

	ErrorInvalidNameRange     = NewError(400, "Name is out of range: (%d - %d)", nameMinLength, nameMaxLength)
	ErrorInvalidEmailRange    = NewError(400, "Email is out of range: (%d - %d)", emailMinLength, emailMaxLength)
	ErrorInvalidEmailPattern  = NewError(400, "The provided email is not in the expected pattern")
	ErrorPasswordCase         = NewError(400, "Passwords must have least one uppercase and lowercase character")
	ErrorPasswordSpecialChar  = NewError(400, "Passwords must have at least one special character")
	ErrorInvalidPasswordRange = NewError(400, "Password is out of range: (%d - %d)", passwordMinLength, passwordMaxLength)

	ErrorIncorrectDateFormat     = NewError(400, "Incorrect date format")
	ErrorDateInThePast           = NewError(400, "Date is in the past")
	ErrorCurrentExchangeNotFound = NewError(404, "Current dollar exchange not found")
)

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewError(code int, msg string, args ...any) *APIError {
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	return &APIError{Code: code, Message: msg}
}

func ErrorInvalidPathParam(path, dataType string) *APIError {
	return NewError(400, fmt.Sprintf(`Invalid path parameter: "%s" should be of type %s`, path, dataType))
}

func ErrorParamNotProvided(name string) *APIError {
	return NewError(400, fmt.Sprintf(`Parameter "%s" is not provided`, name))
}
