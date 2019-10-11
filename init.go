package main

import (
	"log"
	"os"
)

func init() {
	apiKey = getVariableOrFatal("TG_APIKEY")
}

func getVariableOrFatal(varName string) string {
	v := os.Getenv(varName)
	if v == "" {
		log.Fatal(varName, " environment key required")
	}

	return v
}
