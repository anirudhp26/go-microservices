package commons

import (
	"syscall"
)

// EnvString returns the value of the environment variable key if it exists, otherwise it returns the fallback value
func EnvString(key, fallback string) string {
	value, ok := syscall.Getenv(key)
	if !ok {
		return fallback
	}
	return value
}
