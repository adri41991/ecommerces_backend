package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Global config instance
var globalConfig *Config

// SetGlobalConfig sets the global config instance
func SetGlobalConfig(cfg *Config) {
	globalConfig = cfg
}

// GetGlobalConfig returns the global config instance
func GetGlobalConfig() *Config {
	return globalConfig
}

// Config almacena toda la configuración de la aplicación.
// Los campos se rellenan desde variables de entorno.
type Config struct {
	Env               string `env:"ENV,required"`
	Port              string `env:"PORT,required"`
	FirebaseProjectID string `env:"FIREBASE_PROJECT_ID,required"`
}

// Load lee las variables de entorno y las carga en la struct Config.
// Prioriza un archivo .env si existe, lo que es útil para desarrollo local.
func Load() *Config {
	// Carga el archivo .env si existe. Ignora el error si no lo encuentra.
	_ = godotenv.Load()

	cfg := &Config{
		Env:               getEnv("ENV", "development"), // Por defecto, "development"
		Port:              getEnv("PORT", "8080"),       // Por defecto, puerto 8080
		FirebaseProjectID: getEnv("FIREBASE_PROJECT_ID", ""),
	}

	log.Printf("Configuración cargada para el entorno: %s", cfg.Env)

	return cfg
}

// getEnv es una función helper para leer una variable de entorno o devolver un valor por defecto.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
