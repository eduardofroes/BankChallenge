package commons

import (
	"fmt"
	"log"
	"net/http"
)

// CheckError x
func CheckError(err error, message string) {
	if err != nil {
		LogError(err, message)
		panic(fmt.Errorf("%s; error: %s", message, err.Error()))
	}
}

// LogError x
func LogError(err error, message string) {
	if err != nil {
		log.Println("Error:", message, err)
	}
}

//HandleError x
func HandleError(w http.ResponseWriter) {
	r := recover()

	if r != nil {
		err := map[string]string{
			"error": fmt.Sprintf("%s", r),
		}

		WriteJSON(w, err, 500)
	}
}
