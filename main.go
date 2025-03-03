package main

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"

	"Secure-File-Storage/handlers"
	"Secure-File-Storage/middleware"
)

func main() {

	mux := http.NewServeMux()

	mux.Handle("/upload", middleware.Auth(http.HandlerFunc(handlers.UploadFile)))
	mux.Handle("/get", middleware.Auth(http.HandlerFunc(handlers.GetFiles)))
	mux.Handle("/download", middleware.Auth(http.HandlerFunc(handlers.DownloadFile)))
	mux.Handle("/delete", middleware.Auth(http.HandlerFunc(handlers.DeleteFile)))

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "X-Api-Key"},
	})
	handler := c.Handler(mux)

	fmt.Println("Server is running at http://localhost:8000")
	if err := http.ListenAndServe(":8000", handler); err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
}

