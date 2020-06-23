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

	convertByte, err := json.Marshal(mapdata)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(convertByte)

}
