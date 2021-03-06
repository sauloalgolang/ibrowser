package endpoints

import (
	"encoding/json"
	// "go-contacts/models"
	// u "go-contacts/utils"
	// "github.com/gorilla/mux"
	"net/http"
	// "strconv"
)

// https://github.com/adigunhammedolalekan/go-contacts/blob/master/utils/util.go

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
