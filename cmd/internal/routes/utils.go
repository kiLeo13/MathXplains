package routes

import (
	"MathXplains/internal/service"
	"strconv"
)

type R map[string]any

func ToInt(fieldName, in string) (int, *service.APIError) {
	out, err := strconv.Atoi(in)
	if err != nil {
		return 0, service.NewError(400, `Invalid datatype, field "%s" must be of type integer`, fieldName)
	}
	return out, nil
}
