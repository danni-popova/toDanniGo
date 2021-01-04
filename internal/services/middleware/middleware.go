package middleware

import (
	"net/http"
	"strings"

	"github.com/danni-popova/todannigo/internal/pkg/token"

	log "github.com/sirupsen/logrus"
)

type HTTPReqInfo struct {
	// GET etc.
	method    string
	uri       string
	referer   string
	userAgent string
	body      string
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqInfo := &HTTPReqInfo{
			method:    r.Method,
			uri:       r.URL.String(),
			referer:   r.Header.Get("Referer"),
			userAgent: r.Header.Get("User-Agent"),
		}
		log.Println(reqInfo)
		next.ServeHTTP(w, r)
	})
}

// Middleware function, which will be called for each request
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tkn := r.Header.Get("Authorization")
		splitToken := strings.Split(tkn, "Bearer ")

		if len(splitToken) < 2 {
			http.Error(w, "Missing Auth Token", http.StatusUnauthorized)
			return
		}

		tkn = splitToken[1]
		if ctx, ok := token.IsValid(tkn, r.Context()); !ok {
			// Write an error and stop the handler chain
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		} else {
			// We found the token in our map
			log.Printf("Authenticated token %s", tkn)
			// TODO: Instead of the whole token, just send the user ID
			req := r.WithContext(ctx)
			next.ServeHTTP(w, req)
		}
	})
}
