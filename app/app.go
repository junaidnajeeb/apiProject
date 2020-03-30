package app

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"apiProject/controller"
	"apiProject/modal"
	"apiProject/utils"

	jwt "github.com/dgrijalva/jwt-go"

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
	router.HandleFunc("/api/users/login", controller.LoginUserHandler).Methods("POST")

	//Authenticated endpoints
	router.HandleFunc("/api/secret/product", controller.GetSecretProductHandler).Methods("GET")
	router.HandleFunc("/api/secret/project", controller.GetSecretProjectHandler).Methods("GET")

	router.Use(JwtAuthentication) //attach JWT auth middleware

	port := viper.GetString("api.port")

	utils.GetLogger().Info("Starting API on port:", port)

	var connectionUrl = ":" + port

	utils.GetLogger().Fatal(http.ListenAndServe(connectionUrl, router))
}

var JwtAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authRequired := []string{"/api/secret/product", "/api/secret/project"} //List of endpoints that does requires auth
		requestPath := filepath.Clean(r.URL.Path)                              //current request path

		//check
		// I know i should be doing the other way around but this is just a demo
		doAuth := false
		for _, value := range authRequired {
			if value == requestPath {
				doAuth = true
				break
			}
		}
		if doAuth == false {
			next.ServeHTTP(w, r)
			return
		}

		response := make(map[string]interface{})

		//TODO:: check cookie too
		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" { //Token is missing, returns with error code 403 Unauthorized
			response = utils.Message(false, "Missing auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}

		splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
		if len(splitted) != 2 {
			response = utils.Message(false, "Invalid/Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}

		tokenPart := splitted[1] //Grab the token part, what we are truly interested in
		tokenClaim := &modal.TokenClaim{}

		token, err := jwt.ParseWithClaims(tokenPart, tokenClaim, func(token *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString("jwt.appSecret")), nil
		})

		if err != nil { //Malformed token, returns with http code 403 as usual
			response = utils.Message(false, "Malformed authentication token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}

		if !token.Valid { //Token is invalid, maybe not signed on this server
			response = utils.Message(false, "Token is not valid.")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}

		//All good

		ctx := context.WithValue(r.Context(), "user", tokenClaim.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r) //proceed in the middleware chain!
	})
}
