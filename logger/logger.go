package logger

import (
	"log/slog"
	"os"
)

// Env es el tipo para el entorno de la aplicación.
type Env string

const (
	EnvDevelopment Env = "development"
	EnvProduction  Env = "production"
)

// GetEnvType determina el tipo de entorno basado en la cadena de entorno proporcionada.
func GetEnvType(env string) Env {
	switch env {
	case "development":
		return EnvDevelopment
	case "production":
		return EnvProduction
	default:
		return EnvDevelopment
	}
}

// New crea una nueva instancia de slog.Logger según el entorno.
// En desarrollo, usa un formato de texto legible.
// En producción, usa un formato JSON para una mejor integración con sistemas de logs.
func New(env Env) *slog.Logger {
	var log *slog.Logger

	switch env {
	case EnvDevelopment:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case EnvProduction:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default:
		// Por defecto, un logger de producción
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
