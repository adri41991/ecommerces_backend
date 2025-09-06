package pubsub

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"ecommerces/backend/core/http/response"
)

// PubSubMessage represents the format of a message sent by Google Pub/Sub.
type PubSubMessage struct {
	Message struct {
		Attributes  map[string]string `json:"attributes,omitempty"`
		Data        string            `json:"data"`
		MessageID   string            `json:"messageId"`
		PublishTime string            `json:"publishTime"`
	} `json:"message"`
	Subscription string `json:"subscription"`
}

// HandlePubSubEvent is a generic handler for Pub/Sub events.
// It parses the message, calls the provided processor function, and always responds with 200 OK (as expected by Pub/Sub).
// The processor function should handle business logic and return an error if something fails.
func HandlePubSubEvent[T any](w http.ResponseWriter, r *http.Request, log *slog.Logger, processor func(T) error) {
	// Parse Pub/Sub message
	msg, err := parsePubSubMessage[T](r, log)
	if err != nil {
		response.RespondWithError(w, http.StatusBadRequest, "Invalid Pub/Sub message")
		return
	}

	// Process the event
	if err := processor(msg); err != nil {
		log.Error("Event processing failed", slog.Any("error", err))
		// Pub/Sub expects 200 even on error, just log
	}

	// Always respond OK for Pub/Sub
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// parsePubSubMessage parses a Google Cloud Pub/Sub message from an HTTP request.
// It reads the body, unmarshals to PubSubMessage, decodes the base64 data,
// and unmarshals the inner payload to the specified type T.
// Returns the parsed payload of type T or an error.
func parsePubSubMessage[T any](r *http.Request, log *slog.Logger) (T, error) {
	var zero T // Zero value for T

	// Read request body
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error("Error reading request body", slog.Any("error", err))
		return zero, err
	}
	defer r.Body.Close()

	// Log raw body for debugging
	log.Info("Raw Pub/Sub body", slog.String("body", string(bodyBytes)))

	// Unmarshal to PubSubMessage
	var pubsubMsg PubSubMessage
	if err := json.Unmarshal(bodyBytes, &pubsubMsg); err != nil {
		log.Error("Error unmarshalling PubSubMessage", slog.Any("error", err))
		return zero, err
	}

	log.Info("Decoded PubSubMessage", slog.Any("message", pubsubMsg))

	if pubsubMsg.Message.Data == "" {
		log.Error("Data field is empty in Pub/Sub message")
		return zero, errors.New("empty data in Pub/Sub message")
	}

	// Decode base64 data
	decoded, err := base64.StdEncoding.DecodeString(pubsubMsg.Message.Data)
	if err != nil {
		log.Error("Error decoding base64 data", slog.Any("error", err))
		return zero, err
	}

	log.Info("Decoded base64 data", slog.String("data", string(decoded)))

	// Unmarshal inner payload to T
	var result T
	if err := json.Unmarshal(decoded, &result); err != nil {
		log.Error("Error unmarshalling inner payload", slog.Any("error", err))
		return zero, err
	}

	log.Info("Parsed payload", slog.Any("payload", result))

	return result, nil
}
