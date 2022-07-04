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
	Id string `json:"id"`
}

func getHome(w http.ResponseWriter, r *http.Request) {
	id := uuid.New()
	decoder := json.NewDecoder(r.Body)

	var urlToConvert convertUrlStruct

	err := decoder.Decode(&urlToConvert)

	if err != nil {
		panic(err)
	}

	err = qrcode.WriteFile(urlToConvert.Url, qrcode.Medium, 256, "images/qr.png")

	if err != nil {
		panic(err)
	}

	data := &responseQrStruct{
		Id: id.String(),
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
	fs := http.FileServer(http.Dir("images"))

	mux.Handle("/qr/", http.StripPrefix("/qr", fs))
	mux.HandleFunc("/", getHome)

	err := http.ListenAndServe(":3333", mux)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server %s\n", err)

		os.Exit(1)
	}
}
