package helpers

import (
	"log"
	"os"
)

// GetEnv returns an environment variable if it exists or falls back to default if not
func GetEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}

	log.Printf("getEnv: returning fallback environment variable for %s:\"%s\"\n", key, fallback)

	return fallback
}

// SSLCertPath is the path where the ssl cert is located
func SSLCertPath(serviceName string) string {
	// Returns default cert if env var not set
	return GetEnv("SSL_CERT_PATH", "internal/certs/"+serviceName+"/app.crt")
}

// SSLKeyPath is the path where the ssl cert is located
func SSLKeyPath(serviceName string) string {
	// Returns default cert if env var not set
	return GetEnv("SSL_KEY_PATH", "internal/certs/"+serviceName+"/app.key")
}
