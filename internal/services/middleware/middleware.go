package middleware

import (
	"context"
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

	UserInfo userClaims `json:"user_info"`
}

type userClaims struct {
	UserID         int    `json:"user_id"`
	Email          string `json:"email"`
	ProfilePicture string `json:"profile_picture"`
}

// Middleware function, which will be called for each request
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		splitToken := strings.Split(token, "Bearer ")
		token = splitToken[1]

		if ctx, ok := validToken(token, r.Context()); ok {
			// We found the token in our map
			log.Printf("Authenticated token %s", token)
			// TODO: Instead of the whole token, just send the user ID
			req := r.WithContext(ctx)
			next.ServeHTTP(w, req)
		} else {
			// Write an error and stop the handler chain
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

func validToken(tokenString string, ctx context.Context) (context.Context, bool) {
	token, err := jwt.ParseWithClaims(tokenString, &toDanniClaims{}, keyFunc)
	if claims, ok := token.Claims.(*toDanniClaims); ok && token.Valid {
		return context.WithValue(ctx, "user_id", claims.UserInfo.UserID), true
	}

	log.Error(err)
	return ctx, false
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	// TODO: later the "kid" can be used to check the version of the key used to sign the JWT
	// This will come in handy when key rotation is implemented.
	return []byte("the-todanni-secret"), nil
}
