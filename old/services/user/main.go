package main

import (
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"log"
)

func validateAuth(idToken string) string {
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	// Verifies the token is valid and returns it decoded
	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
	}

	uID := token.Claims["user_id"]
	fmt.Println(fmt.Sprintf("%v", uID))

	user, err := client.GetUser(ctx, fmt.Sprintf("%v", uID))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(user.DisplayName)
	return "end"
}

func main(){
	token := "eyJhbGciOiJSUzI1NiIsImtpZCI6ImMzZjI3NjU0MmJmZmU0NWU5OGMyMGQ2MDNlYmUyYmExMTc2ZWRhMzMiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20vdG9kYW5uaS02OTY5IiwiYXVkIjoidG9kYW5uaS02OTY5IiwiYXV0aF90aW1lIjoxNTkyODMyMDAxLCJ1c2VyX2lkIjoiSE9MTEZDdk82bVVhRnJTN3BIdEtUN010ekRuMiIsInN1YiI6IkhPTExGQ3ZPNm1VYUZyUzdwSHRLVDdNdHpEbjIiLCJpYXQiOjE1OTI4MzIwMDEsImV4cCI6MTU5MjgzNTYwMSwiZW1haWwiOiJkYW5uaWkucG9wb3ZhQGdtYWlsLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwiZmlyZWJhc2UiOnsiaWRlbnRpdGllcyI6eyJlbWFpbCI6WyJkYW5uaWkucG9wb3ZhQGdtYWlsLmNvbSJdfSwic2lnbl9pbl9wcm92aWRlciI6InBhc3N3b3JkIn19.bBkWPwmh_fuaSkt79AWTWseeFCBz08WSGOb3qrT3nuW8_3_b9IVKxOMXeYIq6Na-SNa8Ne1ihJqv08HWaN5l8ptfgQVqglww4k6Eerq-ubk1hx1Vsv4_GSMoEtkjKqQOm6h3jRvMM7vRy1XTKwG3WtWUYdzvm8F-LOUX4DY0HSx-yN2ed7RGdLtUoyNAwsODyS0XDKfVDuOEou0jG1hBUSQe4KZ7Ha9ipsG-Pl3ZlD2mxQJzZRzV2_y0I2RX-Bstf8mTGSZUriFcP34G1nC5eOqOT1GQHZ0ifaoyVJ9XNyPKNe-iwwSvxTDxX7vBI1EdtXCWesrhKbiDGcaId3Mkvg"
	fmt.Println(validateAuth(token))
}