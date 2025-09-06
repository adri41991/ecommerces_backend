package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

// responseWriterInterceptor es un wrapper alrededor de http.ResponseWriter
// para capturar el código de estado de la respuesta.
type responseWriterInterceptor struct {
	http.ResponseWriter
	statusCode int
}

func (w *responseWriterInterceptor) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

// NewResponseWriterInterceptor crea una nueva instancia de nuestro response writer.
func NewResponseWriterInterceptor(w http.ResponseWriter) *responseWriterInterceptor {
	// Por defecto, el código de estado es 200 OK.
	return &responseWriterInterceptor{w, http.StatusOK}
}

// Logger es un middleware que registra cada petición HTTP entrante.
func Logger(log *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Usamos nuestro interceptor para capturar el status code
			win := NewResponseWriterInterceptor(w)

			// Pasamos el control al siguiente handler en la cadena
			next.ServeHTTP(win, r)

			log.Info("Request completado",
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.Duration("duration", time.Since(start)),
				slog.Int("status_code", win.statusCode),
			)
		})
	}
}
