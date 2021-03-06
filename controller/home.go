package controller

import (
	"apiProject/modal"
	"encoding/json"
	"net/http"

	"apiProject/utils"

	"github.com/spf13/viper"
)

func HomeLinkHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, applicationJSON)
	utils.LoggerInfo("homeLinkHandler called")

	h := modal.Home{Version: viper.GetString("version"), Message: "Welcome home!"}
	response := utils.Message(true, "Home page")
	response["home"] = h

	json.NewEncoder(w).Encode(response)
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, applicationJSON)
	utils.LoggerInfo("Ping called")

	h := modal.Home{Version: viper.GetString("version"), Message: "Pong"}
	response := utils.Message(true, "Ping page")
	response["ping"] = h

	json.NewEncoder(w).Encode(response)
}
