package webserver

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Server is a wrapper for the Logger and mux router
type Server struct {
	logger *log.Logger
	mux    *http.ServeMux
}

// StartServer Starts the web server on the given port with the given data
func StartServer(port string) {
	router := mux.NewRouter()
	router.HandleFunc("/", HandleHealthCheck)
	listenAddress := fmt.Sprintf(":%s", port)
	server := &http.Server{Addr: listenAddress, Handler: router}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
