package middleware

import (
	"context"
	"net/http"
)

// Key for storing agent UUID in request context
type contextKey string

const AgentUUIDKey contextKey = "agentUUID"

// UUIDMiddleware extracts the agent UUID from headers
func UUIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Extract the UUID from the header
		agentUUID := r.Header.Get("X-Agent-ID")

		// Add the UUID to the request context
		ctx := context.WithValue(r.Context(), AgentUUIDKey, agentUUID)

		// Call the next handler with the updated context
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}
