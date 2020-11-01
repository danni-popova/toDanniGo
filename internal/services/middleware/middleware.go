package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/danni-popova/todannigo/internal/services/claims"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
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
	token, err := jwt.ParseWithClaims(tokenString, &claims.ToDanniClaims{}, keyFunc)
	if err != nil {
		log.Error(err)
		return ctx, false
	}

	if clms, ok := token.Claims.(*claims.ToDanniClaims); ok && token.Valid {
		return context.WithValue(ctx, "user_id", clms.UserInfo.UserID), true
	}

	log.Error(err)
	return ctx, false
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	// TODO: later the "kid" can be used to check the version of the key used to sign the JWT
	// This will come in handy when key rotation is implemented.
	return []byte(claims.HmacSampleSecret), nil
}
