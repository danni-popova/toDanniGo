package claims

import "github.com/dgrijalva/jwt-go"

const (
	// Store this as environment variable in the future
	HmacSampleSecret = "the-todanni-secret"
	// This isn't used anywhere... yet
	TokenIssuer = "todanni-user-service"
)

type ToDanniClaims struct {
	jwt.StandardClaims

	UserInfo UserClaims `json:"user_info"`
}

type UserClaims struct {
	UserID         int    `json:"user_id"`
	Email          string `json:"email"`
	ProfilePicture string `json:"profile_picture"`
}
