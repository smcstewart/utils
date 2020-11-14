package utils

import (
	"os"
	"strings"
)

// EnvOrDefault looks up an environment variable, returning its value or the
// default if not present.
func EnvOrDefault(envVar string, def string) string {
	v, set := os.LookupEnv(envVar)
	if !set {
		v = def
	}
	return v
}

// StrExists performs a case-insensitive, non-sorted lookup (via range) for
// the given key in the slice. Returns true as soon as it finds a match.
func StrExists(slice []string, key string) bool {
	for _, v := range slice {
		if strings.ToLower(v) == key {
			return true
		}
	}
	return false
}
