package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func DecodeRequest(r *http.Request, dst any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(dst)
}

func FormatDecodeError(err error) string {
	var syntaxErr *json.SyntaxError
	var typeErr *json.UnmarshalTypeError

	switch {
	case errors.As(err, &syntaxErr):
		return "invalid JSON syntax"
	case errors.As(err, &typeErr):
		return fmt.Sprintf("%s must be a %s, got %s", typeErr.Field, typeErr.Type, typeErr.Value)
	case strings.HasPrefix(err.Error(), "json: unknown field"):
		field := strings.TrimPrefix(err.Error(), "json: unknown field ")
		return fmt.Sprintf("unknown field %s", field)
	default:
		return "invalid request body"
	}
}
