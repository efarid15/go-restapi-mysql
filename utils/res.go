package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, payload interface{}, status int) {

	mapdata := map[string]interface{}{}
	mapdata["status"] = http.StatusOK
	mapdata["message"] = "Berhasil"
	mapdata["data"] = payload

	response, err := json.Marshal(mapdata)

	maperror := map[string]interface{}{}
	maperror["status"] = http.StatusInternalServerError
	maperror["message"] = "Error"
	maperror["data"] = payload

	responseerror, err := json.Marshal(maperror)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	if status == http.StatusInternalServerError {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, _ = w.Write(responseerror)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, _ = w.Write(response)
	}



}
