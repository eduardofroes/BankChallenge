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

//HandleNotFound function is responsible for send not found http status code.
func HandleNotFound(w http.ResponseWriter, message string) {
	messageBody := map[string]string{
		"message": message,
	}

	WriteJSON(w, messageBody, 404)
}

//HandleUnauthorized function is responsible for send unauthorized http status code.
func HandleUnauthorized(w http.ResponseWriter, message string) {

	body := map[string]string{
		"message": message,
	}

	WriteJSON(w, body, 401)
}

//HandleBadRequest function is responsible for send bad request http status code.
func HandleBadRequest(w http.ResponseWriter, message string) {

	body := map[string]string{
		"message": message,
	}

	WriteJSON(w, body, 400)
}

//HandleNotAcceptable function is responsible for send not acceptable http status code.
func HandleNotAcceptable(w http.ResponseWriter, message string, code string) {

	body := map[string]string{
		"message": message,
		"code":    code,
	}

	WriteJSON(w, body, 406)
}

func TreatCode(w http.ResponseWriter, code string) {

	switch code {
	case "ACCOUNT_001":
	case "ACCOUNT_002":
		HandleBadRequest(w, getMessageCode(code))
		break
	case "FOUNDS_001":
		HandleNotAcceptable(w, getMessageCode(code), code)
		break
	case "TRANFER_001":
		body := map[string]string{
			"message": getMessageCode(code),
		}
		WriteJSON(w, body, 200)
		break
	}
}

func getMessageCode(code string) string {

	codeMap := map[string]string{
		"FUNDS_001":   "Insufients funds.",
		"ACCOUNT_001": "Origin account not found.",
		"ACCOUNT_002": "Destination account not found.",
		"TRANFER_001": "Transfer executed with sucessfuly.",
	}

	return codeMap[code]
}
