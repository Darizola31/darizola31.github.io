package server

import (
	"fmt"
	"github.com/faanross/orlokC2/internal/router"
	"github.com/faanross/orlokC2/internal/websocket"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const serverAddr = "127.0.0.1"
const serverPort = ":8080"

func main() {

	signChan := make(chan os.Signal, 1)
	signal.Notify(signChan, syscall.SIGINT, syscall.SIGTERM)
	websocket.StartWebSocketServer()
	r := chi.NewRouter()

	router.SetupRoutes(r)

	serverAddrPort := fmt.Sprintf("%s:%s", serverAddr, serverPort)

	log.Printf("Server listening on %s\n", serverAddrPort)

	go func() {
		err := http.ListenAndServe(serverAddrPort, r)
		if err != nil {
			log.Fatal(err)
		}
	}()

	<-signChan

	fmt.Println("This will happen!")

}
