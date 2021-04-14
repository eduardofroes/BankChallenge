package commons

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(CPF string, secret string) (*string, error) {

	atClaims := jwt.MapClaims{}

	atClaims["authorized"] = true
	atClaims["cpf"] = CPF
	atClaims["secret"] = secret
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))

	if err != nil {
		return nil, err
	}

	return &token, nil

}

func ExtractToken(autorization string) string {

	atSlit := strings.Split(autorization, " ")

	if len(atSlit) == 2 {
		return atSlit[1]
	}

	return ""
}

func ValidateToken(autorization string) bool {

	if autorization == "" {
		return false
	}

	token := verifyToken(autorization)

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return false
	}

	return true
}

func verifyToken(autorization string) *jwt.Token {
	tokenString := ExtractToken(autorization)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})

	CheckError(err, "Error to validate token.")

	return token
}
