package password

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetHashingCost(hashedPassword []byte) int {
	cost, _ := bcrypt.Cost(hashedPassword)
	return cost
}

func PassWordHashingHandler(w http.ResponseWriter, r *http.Request) {
	password := "secret"
	hash, _ := HashPassword(password)

	fmt.Fprintln(w, "Password:", password)
	fmt.Fprintln(w, "Hash:    ", hash)

	match := CheckPasswordHash(password, hash)
	fmt.Fprintln(w, "Match:    ", match)

	cost := GetHashingCost([]byte(hash))
	fmt.Fprintln(w, "Cost:      ", cost)
}

// func RegisterRoutes(r *mux.Router) {
// 	...
// 	indexRouter := r.PathPrefix("/index").Subrouter()
// 	indexRouter.HandleFunc("/password_hashing", handler.PassWordHashingHandler)
// 	...
// }

//curl -X GET http://localhost:8999/index/password_hashing
// Password: secret
// Hash:     $2a$14$Ael8nW7UF/En/iI7LGdyBuaIO8VREbL2CAShRN0EUQHqtmOHXh.XK
// Match:    true
// Cost:     14
