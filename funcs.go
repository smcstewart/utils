package utils

import "os"

// EnvOrDefault looks up an environment variable, returning its value or the
// default if not present.
func EnvOrDefault(envVar string, def string) string {
	v, set := os.LookupEnv(envVar)
	if !set {
		v = def
	}
	return v
}


