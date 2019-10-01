package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/syuni/pact-sample/pb/addressbook"
	"github.com/syuni/pact-sample/provider/store"
)

// ListPerson provides list of all people in the addressbook
func ListPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	addressbook := store.GetPeople()
	bytes, err := json.MarshalIndent(addressbook, "", "  ")
	if err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Fprint(w, string(bytes))
}
