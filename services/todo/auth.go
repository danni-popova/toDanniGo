package main

import (
	"context"
	firebase "firebase.google.com/go"
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

	return token.UID
}
//
//func login(){
//	ctx := context.Background()
//	app, err := firebase.NewApp(ctx, nil)
//	if err != nil {
//		log.Fatalf("error initializing app: %v\n", err)
//	}
//
//	client, err := app.Auth(ctx)
//	if err != nil {
//		log.Fatalf("error getting Auth client: %v\n", err)
//	}
//}
