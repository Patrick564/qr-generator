package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Patrick564/qr-converter/utils"
	"github.com/google/uuid"
	"github.com/skip2/go-qrcode"
)

type ConvertUrlStruct struct {
	Url string
}

type ResponseQrStruct struct {
	Id       string `json:"id"`
	ImageURL string `json:"image_url"`
}

func CreateNewQR(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := uuid.New()
	imageName := fmt.Sprintf("images/%s.png", id.String())
	imageURL := fmt.Sprintf("%s/api/code/%s.png", r.Host, id.String())

	urlToConvert := ConvertUrlStruct{}
	err := json.NewDecoder(r.Body).Decode(&urlToConvert)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.CustomError{Message: err.Error()})
		return
	}

	// This is commented for future Redis implementation
	// for saving image as array of bytes or bitmap.
	// q, _ := qrcode.Encode(urlToConvert.Url, qrcode.Medium, 256)
	// fmt.Println("qrcode:")
	// fmt.Println(q)

	err = qrcode.WriteFile(
		urlToConvert.Url,
		qrcode.Medium,
		256,
		imageName,
	)
	if err != nil {
		return
	}

	data := &ResponseQrStruct{
		Id:       id.String(),
		ImageURL: imageURL,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(jsonData)
}
