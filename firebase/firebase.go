package firebase

import (
	"context"
	"google.golang.org/api/option"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"log"
)

func InitializeAppWithRefreshToken() (*auth.Client, error) {
	opt := option.WithCredentialsFile("/home/oonya/heroku_apps/memory-share-firebase-adminsdk-zz9zi-4b08f04bd5.json")
	config := &firebase.Config{ProjectID: "memory-share"}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	x, x1 := app.Auth(context.Background())
	return x, x1
}