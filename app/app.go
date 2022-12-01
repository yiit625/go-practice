package app

import (
	"fmt"
	"github.com/ashishjuyal/banking-lib/logger"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func Start() {
	fmt.Println("It is working")
	router := mux.NewRouter()

	// starting server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	logger.Info(fmt.Sprintf("Starting server on %s:%s ...", address, port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}
