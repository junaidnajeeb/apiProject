package controller

import (
	"apiProject/modal"
	"apiProject/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const contentType = "Content-Type"
const applicationJSON = "application/json"

/*
func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, applicationJSON)
	json.NewEncoder(w).Encode(modal.UserList)
}
*/

func UserCreateHandler(w http.ResponseWriter, r *http.Request) {
	utils.LoggerInfo("UserCreateHandler called")
	w.Header().Set(contentType, applicationJSON)

	userPointer := &modal.User{}
	utils.GetLogger().Info(userPointer)
	err := json.NewDecoder(r.Body).Decode(userPointer)

	if err != nil {
		json.NewEncoder(w).Encode(modal.ErrorMessageInstance("Invalid user request"))
		return
	}

	userCreated := userPointer.UserCreate() //Create user

	if userCreated["status"] == false {
		json.NewEncoder(w).Encode(&userCreated)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&userCreated)

}

func GetOneUserHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set(contentType, applicationJSON)
	userID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(modal.ErrorMessageInstance("Bad Requested user Id"))
		return
	}

	tempUser := &modal.User{}
	response := tempUser.GetOneUser(userID)

	json.NewEncoder(w).Encode(response)
	return

}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, applicationJSON)
	userID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(modal.ErrorMessageInstance("Bad Requested user Id"))
		return
	}

	tempUser := &modal.User{}
	response := tempUser.DeleteOneUser(userID)

	json.NewEncoder(w).Encode(response)
	return

}
