package receiverApp

import (
	"fmt"
	"github.com/ashishjuyal/banking-lib/logger"
	"github.com/ashishjuyal/banking/domain"
	"github.com/ashishjuyal/banking/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"
	"time"
)

func Start() {
	fmt.Println("It is working")

	// sanityCheck()

	router := mux.NewRouter()

	// DB Configuration
	dbClient := getDbClient()
	customerRepositoryDb := domain.NewFileRepositoryDb(dbClient)
	ah := FileHandler{service.NewFileService(customerRepositoryDb)}
	router.HandleFunc("/upload", ah.NewImage).Methods(http.MethodPost).Name("UploadFile")
	// starting server
	/*address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	logger.Info(fmt.Sprintf("Starting server on %s:%s ...", address, port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))*/

	port := os.Getenv("PORT")
	logger.Info(fmt.Sprintf("Starting server on %s ...", port))
	log.Fatal(http.ListenAndServe(":"+port, nil), router)
}

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}

func sanityCheck() {
	envProps := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
		"DB_USER",
		"DB_PASSWD",
		"DB_ADDR",
		"DB_PORT",
		"DB_NAME",
	}
	for _, k := range envProps {
		if os.Getenv(k) == "" {
			logger.Fatal(fmt.Sprintf("Environment variable %s not defined. Terminating application...", k))
		}
	}
}
