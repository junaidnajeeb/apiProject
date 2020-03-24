package controller

import (
	"apiProject/modal"
	"encoding/json"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func HomeLinkHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, applicationJSON)
	log.Print("homeLinkHandler called")

	h := modal.Home{Version: viper.GetString("version"), Message: "Welcome home!"}

	json.NewEncoder(w).Encode(h)
}
