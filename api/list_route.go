package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type FilesList struct {
	Files []string `json:"files"`
}

func ListAll(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	files, err := os.ReadDir("./images")
	if err != nil {
		return
	}

	filesList := FilesList{Files: make([]string, 0)}

	for _, f := range files {
		if f.Name() == ".keep" {
			continue
		}
		filesList.Files = append(filesList.Files, fmt.Sprintf("/api/code/%s", f.Name()))
	}

	response, err := json.Marshal(filesList)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
