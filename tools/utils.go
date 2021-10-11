package tools

import (
	"log"
	"os"
	"strings"
)

const SourcePrefix = "SOURCE"

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func LoadEnvConfig() map[string]string {
	config := make(map[string]string)

	for _, e := range os.Environ() {
		if strings.HasPrefix(e, SourcePrefix) {
			pair := strings.SplitN(e, "=", 2)
			item := strings.SplitN(pair[1], ",", 2)
			log.Printf("load config from env %v", item)
			config[item[0]] = item[1]
		}
	}
	return config
}
