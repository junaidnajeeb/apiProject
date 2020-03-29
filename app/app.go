package app

import (
	"fmt"
	"net/http"
	"path/filepath"
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
	router.HandleFunc("/api/", controller.HomeLinkHandler)
	router.HandleFunc("/api/ping", controller.PingHandler)

	// Account endpoints
	router.HandleFunc("/api/accounts", controller.AccountCreateHandler).Methods("POST")
	router.HandleFunc("/api/accounts/{id}", controller.GetOneAccountHandler).Methods("GET")
	router.HandleFunc("/api/accounts/{id}", controller.DeleteAccountHandler).Methods("DELETE")

	// User endpoints
	//router.HandleFunc("/api/users", controller.GetAllUsersHandler).Methods("GET")
	router.HandleFunc("/api/users/{id}", controller.GetOneUserHandler).Methods("GET")
	router.HandleFunc("/api/users", controller.UserCreateHandler).Methods("POST")
	//router.HandleFunc("/api/users", controller.UserUpdateHandler).Methods("PATCH") //TODO:: do this
	router.HandleFunc("/api/users/{id}", controller.DeleteUserHandler).Methods("DELETE")

	router.Use(JwtAuthentication) //attach JWT auth middleware

	port := viper.GetString("api.port")

	utils.GetLogger().Info("Starting API on port:", port)

	var connectionUrl = ":" + port

	utils.GetLogger().Fatal(http.ListenAndServe(connectionUrl, router))
}

var JwtAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		noAuthRequired := []string{"/api", "/api/ping"} //List of endpoints that doesn't require auth
		requestPath := filepath.Clean(r.URL.Path)       //current request path

		//check if request does not need authentication, serve the request if it doesn't need it
		for _, value := range noAuthRequired {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		//TODO:: Implement the logic
		utils.LoggerDebug("TODO:: Implement the logic for " + requestPath)
		next.ServeHTTP(w, r)
		return
	})
}
