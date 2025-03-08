package utils

import (
	"os"
	"strings"
)

func Getenv(key string) *string {
	if key == "" {
		return nil
	} else {
		val := os.Getenv(key)
		if len(val) == 0 {
			return nil
		}
		return &val
	}
}

func GetEnvCors(key string) (fallback []string) {
	value := Getenv(key)
	return strings.Split(*value, ",")
}
