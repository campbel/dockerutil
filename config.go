package dockerutil

import (
	"io/ioutil"
	"os"
	"strings"
)

// Getenv will load the config from 1) environment, 2) Docker secret, 3) Docker config
func Getenv(key string) string {

	// check environment variable
	val := os.Getenv(key)
	if val != "" {
		return val
	}

	// check Docker secret
	out, err := ioutil.ReadFile("/run/secrets/" + key)
	if err == nil {
		return strings.TrimSpace(string(out))
	}

	// check Docker config
	out, err = ioutil.ReadFile("/" + key)
	if err == nil {
		return strings.TrimSpace(string(out))
	}

	return ""
}
