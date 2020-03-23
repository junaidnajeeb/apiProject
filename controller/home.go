package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

func HomeLinkHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, applicationJSON)
	log.Print("homeLinkHandler called")
	json.NewEncoder(w).Encode("Welcome home!")
}
