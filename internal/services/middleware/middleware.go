package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

//func LoggingMiddleware(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		// Do stuff here
//		log.Println(r.RequestURI)
//		// Call the next handler, which can be another middleware in the chain, or the final handler.
//		next.ServeHTTP(w, r)
//	})
//}

type toDanniClaims struct {
	jwt.StandardClaims
}

// Middleware function, which will be called for each request
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		splitToken := strings.Split(token, "Bearer ")
		token = splitToken[1]

		if validToken(token) {
			// We found the token in our map
			log.Printf("Authenticated token %s\n", token)
			// Pass down the request to the next middleware (or final handler)
			next.ServeHTTP(w, r)
		} else {
			// Write an error and stop the handler chain
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

func validToken(tokenString string) bool {
	token, err := jwt.ParseWithClaims(tokenString, &toDanniClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("the-todanni-secret"), nil
	})

	if claims, ok := token.Claims.(*toDanniClaims); ok && token.Valid {
		log.Info(claims.ExpiresAt)
		return true
	}

	log.Error(err)
	return false
}
