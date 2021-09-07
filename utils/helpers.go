package utils

import (
	"fmt"
	"log"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

const defaultSecret = "dWy6_3z4.."

// User type for others to use
type User struct {
		firstname string
		lastname string
		email string
		password string
}

func generateJWT(userData User) (string, error) {
	var err error
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = defaultSecret
	}

	user := userData
	token := jwt.NewWithClaims( jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user,
	})

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil{
		log.Fatal(err)
	}

	return tokenString, nil
}

// func TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		var errorObject models.Error
// 		authHeader := r.Header.Get("Authorization")
// 		bearerToken := strings.Split(authHeader, " ")

// 		if len(bearerToken) == 2 {
// 			authToken := bearerToken[1]

// 			token, error := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
// 				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 					return nil, fmt.Errorf("There was an error")
// 				}

// 				return []byte("secret"), nil
// 			})

// 			if error != nil {
// 				errorObject.Message = error.Error()
// 				RespondWithError(w, http.StatusUnauthorized, errorObject)
// 				return
// 			}

// 			if token.Valid {
// 				next.ServeHTTP(w, r)
// 			} else {
// 				errorObject.Message = error.Error()
// 				RespondWithError(w, http.StatusUnauthorized, errorObject)
// 				return
// 			}
// 		} else {
// 			errorObject.Message = "Invalid token."
// 			RespondWithError(w, http.StatusUnauthorized, errorObject)
// 			return
// 		}
// 	})
// }

// HashAndSalt is used for hashing users passwords
func HashAndSalt(password string) string {

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	pwd := []byte(password)

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

// ComparePasswords is used for hashing and then checking if the password hash is equal to the second params
func ComparePasswords(hashedPwd string, plainPwd []byte) bool {

	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
