package commons

import "net/http"

func Unauthorized(w http.ResponseWriter, message string) {

	body := map[string]string{
		"message": message,
	}

	WriteJSON(w, body, 401)
}
