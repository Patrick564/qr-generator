package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/skip2/go-qrcode"
)

type convertUrlStruct struct {
	Url string
}

type responseQrStruct struct {
	Id       string `json:"id"`
	ImageURL string `json:"image_url"`
}

func createNewQR(w http.ResponseWriter, r *http.Request) {
	id := uuid.New()
	decoder := json.NewDecoder(r.Body)
	imageName := fmt.Sprintf("images/%s.png", id.String())

	var urlToConvert convertUrlStruct

	err := decoder.Decode(&urlToConvert)

	if err != nil {
		panic(err)
	}

	err = qrcode.WriteFile(
		urlToConvert.Url,
		qrcode.Medium,
		256,
		imageName,
	)

	if err != nil {
		panic(err)
	}

	data := &responseQrStruct{
		Id:       id.String(),
		ImageURL: fmt.Sprintf("%s/%s.png", r.Host, id.String()),
	}
	jsonData, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func main() {
	mux := http.NewServeMux()
	port := os.Getenv("PORT")
	fs := http.FileServer(http.Dir("images"))

	mux.Handle("/qr-codes/", http.StripPrefix("/qr-codes", fs))
	mux.HandleFunc("/", createNewQR)

	if port == "" {
		port = ":9000"
	}

	err := http.ListenAndServe(port, mux)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server %s\n", err)

		os.Exit(1)
	}
}
