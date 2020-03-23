package controller

import (
	"apiProject/modal"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const contentType = "Content-Type"
const applicationJSON = "application/json"

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, applicationJSON)
	json.NewEncoder(w).Encode(modal.UserList)
}

func GetOneUserHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set(contentType, applicationJSON)
	userID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(modal.APIErrorMessageInstance("Bad Requested user Id"))
		return
	}

	if singleUser, ok := modal.UserList[userID]; ok {
		json.NewEncoder(w).Encode(singleUser)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(modal.APIErrorMessageInstance("Requested user not found"))
	return
}

func CreateUpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser modal.User

	w.Header().Set(contentType, applicationJSON)

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(modal.APIErrorMessageInstance("Bad request"))
		return
	}

	json.Unmarshal(reqBody, &newUser)
	modal.UserAdd(newUser)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newUser)

}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, applicationJSON)

	userID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(modal.APIErrorMessageInstance("Bad Requested user Id"))
		return
	}

	_, ok := modal.UserList[userID]
	if ok {
		delete(modal.UserList, userID)
		json.NewEncoder(w).Encode("Deleted")
		return
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(modal.APIErrorMessageInstance("Requested user not found"))
	return

}
