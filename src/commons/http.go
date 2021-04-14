package commons

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// WriteJSON function is responsible for convert an object to json.
func WriteJSON(w http.ResponseWriter, body interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if body != nil {
		bytes, err := json.Marshal(body)
		CheckError(err, fmt.Sprintf("Error in json serialization process: %#v", body))

		w.Write(bytes)
	}
}
