package respond

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status int `json:"status"`
	Message string `json:"message"`
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	mapdata := map[string]interface{}{}
	mapdata["status"] = http.StatusOK
	mapdata["message"] = "Berhasil"
	mapdata["data"] = payload

	response, err := json.Marshal(mapdata)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
