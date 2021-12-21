package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"exchange-rates-cbr/handlers"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var port string
var allowedOrigin string

func init() {
	godotenv.Load()
	port = fmt.Sprintf(":%s", os.Getenv("PORT"))
	allowedOrigin = os.Getenv("ALLOWED_ORIGIN")
}

func main() {
	sm := mux.NewRouter()

	apiHandlers := handlers.NewApiHandlers()
	apiRouter := sm.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/rates", apiHandlers.Index).Queries("date", "{date}").Methods(http.MethodGet)

	cors := gohandlers.CORS(
		gohandlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		gohandlers.AllowedOrigins([]string{allowedOrigin}),
		gohandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
	)

	server := &http.Server{
		Addr:         port,
		Handler:      cors(sm),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		log.Println("Starting http server at", port)
		err := server.ListenAndServe()
		if err != nil {
			log.Println("Error", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	sig := <-sigChan
	log.Printf("Recieved terminate signal, graceful shutdown, signal: [%s]", sig)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	server.Shutdown(ctx)
}
