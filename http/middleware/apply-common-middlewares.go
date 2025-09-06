package middleware

import (
	"net/http"

	"github.com/adri41991/ecommerces_backend/core/logger"
)

// ApplyCommonMiddlewares aplica middlewares comunes (Logger y Recovery) a cualquier handler.
// Requiere especificar el entorno explícitamente.
func ApplyCommonMiddlewares(handler http.Handler, env logger.Env) http.Handler {
	// Crear logger con el entorno seleccionado
	log := logger.New(env)

	// Aplicar Recovery primero (orden importa: Recovery envuelve Logger)
	handler = Recovery(log)(handler)
	// Luego Logger
	handler = Logger(log)(handler)
	return handler
}
