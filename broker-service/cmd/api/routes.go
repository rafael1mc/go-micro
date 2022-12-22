package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	// specify who is allowed to connect
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	mux.Use(middleware.Heartbeat("/ping"))

	mux.Use(loggerMiddleware)

	mux.Post("/", app.Broker)

	mux.Get("/", app.Broker)

	mux.Post("/handle", app.handleSubmission)

	mux.Post("/log-grpc", app.logItemViaGRPC)

	return mux
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var bodyData string
		if r.Body != nil {
			bodyBytes, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Printf("Error reading body: %v", err)
				return
			}
			bodyData = string(bodyBytes)
			r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		log.Printf("\n%s %s %s %s", r.Method, r.RequestURI, r.Host, bodyData)
		next.ServeHTTP(w, r)
	})
}
