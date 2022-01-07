package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/oonyaTeam/memory_share_server/firebase"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type FirebaseCustomToken struct {
	Kind         string `json:"kind"`
	IDToken      string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
	IsNewUser    bool   `json:"isNewUser"`
}

func main() {
	uid := "uid"
	tokenFileName := ".idToken"

	client, err := firebase.InitializeAppWithRefreshToken()
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	token, err := client.CustomToken(context.Background(), uid)
	if err != nil {
		log.Fatalf("error minting custom token: %v\n", err)
	}

	webapikey := os.Getenv("FS_WEB_API_KEY")
	if webapikey == "" {
		log.Fatal("FIREBASE_WEB_API_KEY is missing")
	}
	endpoint := fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:signInWithCustomToken?key=%s", webapikey)

	body := []byte(fmt.Sprintf(`
	{
		"token":"%s",
		"returnSecureToken":true
	}
	`, token))
	values := url.Values{}
	values.Set("returnSecureToken", "true")
	values.Set("token", token)
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))
	if err != nil {
		log.Fatalf("error minting custom token: %v\n", err)
	}

	req.Header.Set("Content-Type", "application/json")
	c := &http.Client{}
	resp, err := c.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			log.Fatal(closeErr)
		}
	}()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	firebaseCustomToken := &FirebaseCustomToken{}
	if err := json.Unmarshal(respBytes, firebaseCustomToken); err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile(tokenFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if fileCloseErr := file.Close(); fileCloseErr != nil {
			log.Fatal(fileCloseErr)
		}
	}()

	_, err = fmt.Fprint(file, firebaseCustomToken.IDToken)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(respBytes))
	fmt.Printf("%s created", tokenFileName)
}