package controller

import (
	"apiProject/modal"
	"apiProject/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AccountCreateHandler(w http.ResponseWriter, r *http.Request) {
	utils.LoggerInfo("AccountCreateHandler called")
	w.Header().Set(contentType, applicationJSON)

	accountPointer := &modal.Account{}
	utils.GetLogger().Info(accountPointer)
	err := json.NewDecoder(r.Body).Decode(accountPointer)

	if err != nil {
		json.NewEncoder(w).Encode(modal.ErrorMessageInstance("Invalid account request"))
		return
	}

	accountCreated := accountPointer.Create() //Create account

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(&accountCreated)

}
func GetOneAccountHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set(contentType, applicationJSON)
	accountID, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(modal.ErrorMessageInstance("Bad Requested user Id"))
		return
	}

	tempAccount := &modal.Account{}
	response := tempAccount.GetOneAccount(accountID)

	json.NewEncoder(w).Encode(response)
	return

}
