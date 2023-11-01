package utils

import (
	"os"
	"strconv"
)

func GetEnv(key string, def string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	return value
}

func GetEnvAsInt(key string, def int) int {
	value, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	v, err := strconv.Atoi(value)
	if err != nil {
		return def
	}
	return v
}
