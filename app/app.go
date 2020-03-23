package app

import (
	"fmt"
	"log"
	"net/http"

	"apiProject/controller"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

// SetupConfiguration Export all the api routes

func SetupConfiguration() {
	fmt.Println("app => SetupConfiguration")

	// name of config file (without extension)
	viper.SetConfigName("api")
	// look for config in the config working directory
	viper.AddConfigPath("config")

	// Find and read the config file
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
}

func SetupRoutes() {
	fmt.Println("app => SetupRoutes")

	// creates a new instance of a mux router
	router := mux.NewRouter().StrictSlash(true)

	// replace http.HandleFunc with myRouter.HandleFunc
	router.HandleFunc("/", controller.HomeLinkHandler)

	// User Map endpoints
	router.HandleFunc("/users", controller.GetAllUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id}", controller.GetOneUserHandler).Methods("GET")
	router.HandleFunc("/users", controller.CreateUpdateUserHandler).Methods("POST")
	router.HandleFunc("/users/{id}", controller.DeleteUserHandler).Methods("DELETE")

	port := viper.GetString("api.port")

	fmt.Println("Starting API on port:", port)

	var connectionUrl = ":" + port

	log.Fatal(http.ListenAndServe(connectionUrl, router))
}
