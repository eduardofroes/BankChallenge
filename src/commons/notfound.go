package commons

import "net/http"

//HandleNotFound x
func HandleNotFound(w http.ResponseWriter, message string) {
	messageBody := map[string]string{
		"message": message,
	}

	WriteJSON(w, messageBody, 404)
}
