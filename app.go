package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	serverPort = "80"
)

var logger *log.Logger
var accounts []Account

func main() {
	InitializeLogging()
	ConnectDB()

	// temporary test data
	// accounts = append(accounts, Account{ID: "1", Firstname: "Anthony", Lastname: "Shewnarain", EmailAddress: "anthony.shewnarain@gmail.com"})

	router := mux.NewRouter()
	router.HandleFunc("/", SayHello).Methods("GET")
	InitializeEndpoints(router)

	// start the server
	done := make(chan bool)
	go func() {
		log.Fatal(http.ListenAndServe(":"+serverPort, router))
	}()
	logger.Printf("Baigan Service started at port %v...", serverPort)
	<-done
}

// InitializeLogging initialize the logger
func InitializeLogging() {
	directory := "logs"
	fileName := "baigan-service.log"
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		err = os.MkdirAll(directory, 0755)
		if err != nil {
			panic(err)
		}
	}
	f, err := os.OpenFile(directory+"/"+fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatal(err)
	}
	logger = log.New(f, "", log.Ldate|log.Ltime)
	logger.SetOutput(io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   directory + "/" + fileName,
		MaxSize:    10, // megabytes
		MaxBackups: 1,
		MaxAge:     14,   //days
		Compress:   true, // disabled by default
	}))
}

// InitializeEndpoints initialize all endpoints
func InitializeEndpoints(r *mux.Router) {
	initializeAccountEndpoints(r)
}

// InitializeAccountEndpoints initialize accounts resources
func initializeAccountEndpoints(r *mux.Router) {
	r.HandleFunc("/accounts", GetAccounts).Methods("GET")
	r.HandleFunc("/accounts/{id}", GetAccount).Methods("GET")
	r.HandleFunc("/accounts/{id}", CreateAccount).Methods("POST")
	r.HandleFunc("/accounts/{id}", UpdateAccount).Methods("PUT")
	r.HandleFunc("/accounts/{id}", DeleteAccount).Methods("DELETE")
	return
}

// SayHello Simple ping
func SayHello(w http.ResponseWriter, r *http.Request) {
	logger.Println("HTTP GET /")
	data, err := Asset("data/notes.txt")
	if err != nil {
		fmt.Print(err)
	}
	notes := string(data)
	fmt.Fprintln(w, notes)
}
