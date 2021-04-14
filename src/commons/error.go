package commons

import (
	"fmt"
	"log"
	"net/http"
)

// CheckError is responsible to throw errors.
func CheckError(err error, message string) {
	if err != nil {
		LogError(err, message)
		panic(fmt.Errorf("%s; error: %s", message, err.Error()))
	}
}

// LogError is responsible to log errors.
func LogError(err error, message string) {
	if err != nil {
		log.Println("Error:", message, err)
	}
}

//HandleError is responsible to handle error http to the client.
func HandleError(w http.ResponseWriter) {
	r := recover()

	if r != nil {
		err := map[string]string{
			"error": fmt.Sprintf("%s", r),
		}

		WriteJSON(w, err, 500)
	}
}
