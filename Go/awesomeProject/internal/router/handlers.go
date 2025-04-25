package router

import (
	"github.com/faanross/orlokC2/internal/middleware"
	"log"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	agentUUID, _ := r.Context().Value(middleware.AgentUUIDKey).(string)

	log.Printf("you hit the endpoint%s\n", r.URL.Path)

	log.Printf("Endpoint %s hit by agent %s\n", r.URL.Path, agentUUID)

	w.Write([]byte("I'm Mister Derp!"))
}
