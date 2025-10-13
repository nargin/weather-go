package main

import (
	"log"
	"os"
	"strings"
)

func parseEnv() {
	env, err := os.ReadFile(".env")
	if err != nil {
		log.Fatalf("Error reading .env file: %v", err)
	}

	lines := string(env)
	for _, line := range strings.Split(lines, "\n") {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := parts[0]
			value := parts[1]
			os.Setenv(key, value)
		}
	}
}
