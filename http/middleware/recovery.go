package middleware

import (
	"log/slog"
	"net/http"
	"runtime/debug"

	"ecommerces/backend/core/http/response"
)

// Recovery es un middleware que se recupera de panics, los registra y devuelve
// una respuesta de error 500 Internal Server Error.
func Recovery(log *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// Usamos defer para asegurarnos de que esta función se ejecute
			// justo antes de que la función del handler retorne, incluso si hay un panic.
			defer func() {
				if err := recover(); err != nil {
					// Si recover() devuelve algo, es que hubo un panic.
					log.Error(
						"Panic recuperado",
						slog.Any("error", err),
						// Capturamos y registramos el stack trace para facilitar la depuración.
						slog.String("stack", string(debug.Stack())),
					)

					// Devolvemos una respuesta de error genérica al cliente.
					response.RespondWithError(w, http.StatusInternalServerError, "ERR_INTERNAL_SERVER")
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}
