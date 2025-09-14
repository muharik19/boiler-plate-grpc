package utils

import (
	"net/http"
)

func ConvertStatusResponseCode(responseCode string) uint32 {
	var code uint32
	switch {
	case responseCode == SUCCESS:
		code = http.StatusOK
	case responseCode == FAILED_INTERNAL:
		code = http.StatusInternalServerError
	case responseCode == FAILED_NOT_FOUND:
		code = http.StatusNotFound
	case responseCode == FAILED_REQUIRED:
		code = http.StatusBadRequest
	case responseCode == FAILED_AUTHORIZED:
		code = http.StatusUnauthorized
	case responseCode == FAILED_EXIST:
		code = http.StatusBadRequest
	default:
		code = http.StatusInternalServerError
	}
	return code
}
