package helpers

import "os"

func EnvVarOrFallback(key string, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}

	return value
}
