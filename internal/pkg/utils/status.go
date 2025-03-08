package utils

import (
	"net/http"

	"github.com/muharik19/boiler-plate-grpc/internal/constant"
)

func ConvertStatusResponseCode(responseCode string) uint32 {
	var code uint32
	switch {
	case responseCode == constant.SUCCESS:
		code = http.StatusOK
	case responseCode == constant.FAILED_INTERNAL:
		code = http.StatusInternalServerError
	case responseCode == constant.FAILED_NOT_FOUND:
		code = http.StatusNotFound
	case responseCode == constant.FAILED_REQUIRED:
		code = http.StatusBadRequest
	case responseCode == constant.FAILED_AUTHORIZED:
		code = http.StatusUnauthorized
	case responseCode == constant.FAILED_EXIST:
		code = http.StatusBadRequest
	default:
		code = http.StatusInternalServerError
	}
	return code
}
