package respond

import (
	"encoding/json"
	"net/http"
)


// respond json error
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

// respond json success
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
