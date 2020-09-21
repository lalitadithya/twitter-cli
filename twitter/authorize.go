package twitter

import (
	"fmt"

	"github.com/dghubble/oauth1"
)

func GetAuthorizationURL(client *TwitterClient) (string, error) {
	fmt.Println(client.userSecrets.GetConsumerKey())
	fmt.Println(client.userSecrets.GetConsumerSecret())

	client.config = oauth1.Config{
		ConsumerKey:    client.userSecrets.GetConsumerKey(),
		ConsumerSecret: client.userSecrets.GetConsumerSecret(),
		CallbackURL:    "oob",
		Endpoint:       TwitterAuthenticateEndpoint,
	}

	requestToken, requestSecret, err := client.config.RequestToken()
	if err != nil {
		fmt.Printf("ERROR: %s\n\n", err.Error())
		return "", err
	}
	fmt.Println("Got: ", requestToken, requestSecret)

	authorizationURL, err := client.config.AuthorizationURL(requestToken)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return "", err
	}
	fmt.Println("Got: ", authorizationURL)

	client.requestToken = requestToken
	client.requestSecret = requestSecret
	return authorizationURL.String(), nil
}

func FetchAndSaveOAuthToken(client *TwitterClient, pin string) error {
	accessToken, accessSecret, err := client.config.AccessToken(client.requestToken, client.requestSecret, pin)
	if err != nil {
		return err
	}

	err = client.userSecrets.SetAccessSecrets(accessToken, accessSecret)
	if err != nil {
		return err
	}

	client.token = oauth1.NewToken(accessToken, accessSecret)
	fmt.Println("Got token:", client.token.Token, client.token.TokenSecret)
	return nil
}
