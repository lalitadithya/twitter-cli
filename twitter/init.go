package twitter

import (
	"net/http"

	"github.com/dghubble/oauth1"
	"github.com/lalitadithya/twitter-cli/twitter/util"
)

var userSecrets *util.UserSecrets

func LoadSecrets() (*TwitterClient, error) {
	twitterClient := &TwitterClient{
		userSecrets: &util.UserSecrets{},
	}
	loadErr := twitterClient.userSecrets.LoadConsumerSecrets()
	if loadErr != nil {
		return nil, loadErr
	}

	return twitterClient, nil
}

func SetSecrets(apiKey, apiSecretKey string) (*TwitterClient, error) {
	twitterClient := &TwitterClient{
		userSecrets: &util.UserSecrets{},
	}
	setErr := twitterClient.userSecrets.SetConsumerSecrets(apiKey, apiSecretKey)
	if setErr != nil {
		return nil, setErr
	}

	return twitterClient, nil
}

func GetAuthorizedClient() (*TwitterClient, error) {
	twitterClient := &TwitterClient{
		userSecrets: &util.UserSecrets{},
	}

	loadErr := twitterClient.userSecrets.LoadAccessSecrets()
	if loadErr != nil {
		return nil, loadErr
	}

	loadErr = twitterClient.userSecrets.LoadConsumerSecrets()
	if loadErr != nil {
		return nil, loadErr
	}

	return twitterClient, nil
}

func getHttpClient(client *TwitterClient) *http.Client {
	config := oauth1.NewConfig(client.userSecrets.GetConsumerKey(), client.userSecrets.GetConsumerSecret())
	token := oauth1.NewToken(client.userSecrets.GetAccessToken(), client.userSecrets.GetAccessSecret())
	return config.Client(oauth1.NoContext, token)
}
