package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Patrick564/qr-converter/api"
)

func RoutesLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Before request

		h.ServeHTTP(w, r)

		// After request
		log.Printf("%s - %s", r.Method, r.RequestURI)
	})
}

func main() {
	mux := http.NewServeMux()
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	fileSrv := http.FileServer(http.Dir("images"))

	mux.Handle("/api/code/", http.StripPrefix("/api/code", fileSrv))
	mux.HandleFunc("/api/codes", api.ListAll)
	mux.HandleFunc("/api/create", api.CreateNewQR)

	log.Printf("Server start at port %s\n", port)
	err := http.ListenAndServe(port, RoutesLogger(mux))
	if err != nil {
		log.Print("Server closed, bye!\n")
		os.Exit(1)
	}
}
