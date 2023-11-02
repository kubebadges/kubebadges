package utils

import (
	"os"
	"strconv"
	"testing"
)

func TestGetEnv(t *testing.T) {
	const envKey = "TEST_ENV"
	const defaultValue = "default"

	os.Unsetenv(envKey)
	if v := GetEnv(envKey, defaultValue); v != defaultValue {
		t.Errorf("Expected %q, but got %q", defaultValue, v)
	}

	expectedValue := "value"
	os.Setenv(envKey, expectedValue)
	if v := GetEnv(envKey, defaultValue); v != expectedValue {
		t.Errorf("Expected %q, but got %q", expectedValue, v)
	}
}

func TestGetEnvAsInt(t *testing.T) {
	const envKey = "TEST_ENV"
	const defaultValue = 42

	os.Unsetenv(envKey)
	if v := GetEnvAsInt(envKey, defaultValue); v != defaultValue {
		t.Errorf("Expected %d, but got %d", defaultValue, v)
	}

	os.Setenv(envKey, "not_a_number")
	if v := GetEnvAsInt(envKey, defaultValue); v != defaultValue {
		t.Errorf("Expected %d, but got %d", defaultValue, v)
	}

	expectedValue := 123
	os.Setenv(envKey, strconv.Itoa(expectedValue))
	if v := GetEnvAsInt(envKey, defaultValue); v != expectedValue {
		t.Errorf("Expected %d, but got %d", expectedValue, v)
	}
}
