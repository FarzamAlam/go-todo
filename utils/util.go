package utils

import (
	"encoding/json"
	"net/http"
)

func Respond(w http.ResponseWriter, status int, message interface{}) {
	response, err := json.Marshal(message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(response))
}

// In case of any error we need to call RespondError
func RespondError(w http.ResponseWriter, status int, message string) {
	Respond(w, status, map[string]string{"error": message})
}
