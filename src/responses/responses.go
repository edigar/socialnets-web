package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorAPI struct {
	Error string `json:"error"`
}

func JSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if statusCode != http.StatusNoContent {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}
}

func HandleErrorStatusCode(w http.ResponseWriter, r *http.Response) {
	var responseError ErrorAPI
	err := json.NewDecoder(r.Body).Decode(&responseError)
	if err != nil && r.StatusCode != 404 {
		log.Fatal(err)
	}

	JSON(w, r.StatusCode, responseError)
}
