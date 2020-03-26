package app

import (
	"fmt"
	"net/http"
	"strings"

	"apiProject/controller"
	"apiProject/utils"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

// SetupConfiguration Export all the api routes

func SetupConfiguration() {
	utils.LoggerInfo("app => SetupConfiguration")

	// name of config file (without extension)
	viper.SetConfigName("app")
	// look for config in the config working directory
	viper.AddConfigPath("config")

	// this way we can pass env var and read it as APP_PORT
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	// Find and read the config file
	err := viper.ReadInConfig()

	if err != nil {
		s := fmt.Sprintf("Error while reading config file %s", err)
		utils.LoggerError(s)
	}
}

func SetupRoutes() {

	utils.LoggerInfo("app => SetupRoutes")

	// creates a new instance of a mux router
	router := mux.NewRouter().StrictSlash(true)

	// replace http.HandleFunc with myRouter.HandleFunc
	router.HandleFunc("/", controller.HomeLinkHandler)

	// User endpoints
	router.HandleFunc("/users", controller.GetAllUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id}", controller.GetOneUserHandler).Methods("GET")
	router.HandleFunc("/users", controller.CreateUpdateUserHandler).Methods("POST")
	router.HandleFunc("/users/{id}", controller.DeleteUserHandler).Methods("DELETE")

	// Account endpoints
	router.HandleFunc("/accounts", controller.AccountCreateHandler).Methods("POST")
	router.HandleFunc("/accounts/{id}", controller.GetOneAccountHandler).Methods("GET")

	port := viper.GetString("api.port")

	utils.GetLogger().Info("Starting API on port:", port)

	var connectionUrl = ":" + port

	utils.GetLogger().Fatal(http.ListenAndServe(connectionUrl, router))
}
