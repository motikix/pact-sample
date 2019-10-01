package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/syuni/pact-sample/provider/store"
)

func GetPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	person, err := store.FetchByID(int32(id))
	if err != nil {
		ResponseError(w, http.StatusNotFound, err.Error())
		return
	}
	bytes, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Fprint(w, string(bytes))
}
