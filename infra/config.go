package infra

import (
	"log"
	"os"
)

func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Printf("Variável de ambiente %s não definida. Usando valor padrão: %s", key, defaultValue)
		return defaultValue
	}
	return value
}
