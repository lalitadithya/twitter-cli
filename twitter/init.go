package twitter

import (
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
