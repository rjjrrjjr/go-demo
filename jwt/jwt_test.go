package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"testing"
)

func TestGenerateJWT(t *testing.T) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	//claims["authorized"] = true
	claims["user"] = "user@example.com"
	//claims["exp"] = time.Now().Add(time.Second * 1).Unix()

	tokenString, err := token.SignedString([]byte("secret"))

	fmt.Println(tokenString, err)
}

func TestParseJWT(t *testing.T) {
	tokenString := ""
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})

	if err != nil {
		fmt.Println(err)
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	fmt.Println(claims, ok, token.Valid)

}