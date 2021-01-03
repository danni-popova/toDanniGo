package token

import (
	"context"
	"time"

	"github.com/danni-popova/todannigo/internal/repositories/account"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

const (
	// Store this as environment variable in the future
	Secret = "the-todanni-secret"
	// This isn't used anywhere... yet
	Issuer = "todanni-user-service"
	// Encryption cost
	Cost = 14
)

type ToDanniClaims struct {
	jwt.StandardClaims
	UserInfo UserClaims `json:"user_info"`
}

type UserClaims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
}

func IsValid(tokenString string, ctx context.Context) (context.Context, bool) {
	token, err := jwt.ParseWithClaims(tokenString, &ToDanniClaims{}, keyFunc)
	if err != nil {
		log.Error(err)
		return ctx, false
	}

	if clms, ok := token.Claims.(*ToDanniClaims); ok && token.Valid {
		return context.WithValue(ctx, "user_id", clms.UserInfo.UserID), true
	}

	log.Error(err)
	return ctx, false
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	// TODO: later the "kid" can be used to check the version of the key used to sign the JWT
	// This will come in handy when key rotation is implemented.
	return []byte(Secret), nil
}

func Generate(authDetails account.AuthDetails) string {
	// Create the Claims
	userInfoClaims := UserClaims{
		UserID: authDetails.ID,
		Email:  authDetails.Email,
	}

	clms := ToDanniClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    Issuer,
			ExpiresAt: time.Now().Add(time.Hour * 6).Unix(),
		},
		UserInfo: userInfoClaims,
	}

	// Generate the Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clms)
	ss, err := token.SignedString([]byte(Secret))
	if err != nil {
		log.Error(err)
	}
	return ss
}
