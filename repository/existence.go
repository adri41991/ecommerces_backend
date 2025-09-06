package repository

import (
	"log/slog"
)

// ExistenceChecker es una interfaz para repositorios que pueden verificar existencia.
type ExistenceChecker interface {
	Exists(id string) (bool, error)
}

// CheckExistence verifica si una entidad existe usando el repositorio proporcionado.
// Registra logs y maneja errores de forma consistente.
// log: Logger a usar para registrar errores (inyectado para flexibilidad).
func CheckExistence(checker ExistenceChecker, id string, log *slog.Logger) (bool, error) {
	exists, err := checker.Exists(id)
	if err != nil {
		log.Error("Error checking existence", "id", id, "error", err)
		return false, err
	}
	return exists, nil
}
