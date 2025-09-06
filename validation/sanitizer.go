package validation

import (
	"strings"
)

// TrimAndLower sanitiza un string: elimina espacios en blanco y convierte a minúsculas.
// Útil para campos como roles, emails, etc.
func TrimAndLower(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}
