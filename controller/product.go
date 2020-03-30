package controller

import (
	"encoding/json"
	"net/http"
)

func GetSecretProductHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set(contentType, applicationJSON)
	json.NewEncoder(w).Encode("You saw the secret Product!!!")
}

func GetSecretProjectHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set(contentType, applicationJSON)
	json.NewEncoder(w).Encode("You saw the secret Project!!!")
}
