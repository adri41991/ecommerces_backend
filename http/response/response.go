package response

import (
	"encoding/json"
	"net/http"
)

// RespondWithJSON escribe una respuesta JSON estándar.
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if payload != nil {
		json.NewEncoder(w).Encode(payload)
	}
}

// RespondWithError escribe una respuesta de error JSON estandarizada.
// El payload es un mapa que contiene solo el código de error.
func RespondWithError(w http.ResponseWriter, code int, errCode string) {
	payload := map[string]string{"code": errCode}
	RespondWithJSON(w, code, payload)
}
