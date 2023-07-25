package config

import (
	"fmt"
	"os"
)

var Server = "127.0.0.1:8888"

func getMode() string {
	mode := os.Getenv("GGP_MODE")

	if mode == "" {
		mode = "test"
	}

	return mode
}

func isProduction() bool {
	result := false

	mode := os.Getenv("GGP_MODE")

	if mode == "release" {
		result = true
	}

	return result
}

func getKey(k string) string {
	mode := getMode()
	result := fmt.Sprintf("%s-%s", k, mode)

	return result
}
