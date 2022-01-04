package firebase

import (
	"context"
	"google.golang.org/api/option"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"log"
	"os"
	"fmt"
)

func InitializeAppWithRefreshToken() (*auth.Client, error) {
	opt := option.WithCredentialsFile("./firebase_credentials.json")
	config := &firebase.Config{ProjectID: "memory-share"}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	x, x1 := app.Auth(context.Background())
	return x, x1
}

func CreateFirebaseJson() {
	fp, err := os.Create("./firebase_credentials.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()

	file := fmt.Sprintf(` {
	"type": "%s",
	"project_id": "%s",
	"private_key_id": "%s",
	"private_key": "%s",
	"client_email": "%s",
	"client_id": "%s",
	"auth_uri": "%s",
	"token_uri": "%s",
	"auth_provider_x509_cert_url": "%s",
	"client_x509_cert_url": "%s"
}`,
		os.Getenv("FS_TYPE"),
		os.Getenv("FS_PROJECT_ID"),
		os.Getenv("FS_PRIVATE_KEY_ID"),
		os.Getenv("FS_PRIVATE_KEY"),
		os.Getenv("FS_CLIENT_EMAIL"),
		os.Getenv("FS_CLIENT_ID"),
		os.Getenv("FS_AUTH_URI"),
		os.Getenv("FS_TOKEN_URI"),
		os.Getenv("FS_AUTH_PROVIDER_X509_CERT_URL"),
		os.Getenv("FS_CLIENT_X509_CERT_URL"))

	_, err = fp.Write(([]byte)(file))
	if err != nil {
		fmt.Println(err)
	}
}