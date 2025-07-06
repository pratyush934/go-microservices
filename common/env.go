package common

import "os"

func EnvString(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback

}
