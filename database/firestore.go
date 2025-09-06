package database

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
)

// NewFirestoreClient crea y devuelve un nuevo cliente de Firestore.
// La configuración de credenciales se maneja automáticamente por el SDK de Google Cloud
// a través de variables de entorno o roles de IAM.
func NewFirestoreClient(ctx context.Context, projectID string, databaseID string) (*firestore.Client, error) {
	client, err := firestore.NewClientWithDatabase(ctx, projectID, databaseID)
	if err != nil {
		return nil, fmt.Errorf("error inicializando el cliente de Firestore: %w", err)
	}

	return client, nil
}
