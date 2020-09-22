package util

import (
	"errors"

	"github.com/spf13/viper"
)

type UserSecrets struct {
	apiKey       string
	apiSecretKey string
	accessToken  string
	accessSecret string
}

var APIKeyMissingError = errors.New("API Key Missing")

func (us *UserSecrets) LoadConsumerSecrets() error {
	us.apiKey = viper.GetString("ConsumerAPIKey")
	if len(us.apiKey) == 0 {
		return APIKeyMissingError
	}

	us.apiSecretKey = viper.GetString("ConsumerAPISecretKey")
	if len(us.apiSecretKey) == 0 {
		return APIKeyMissingError
	}

	return nil
}

func (us *UserSecrets) SetConsumerSecrets(apiKey, apiSecretKey string) error {
	us.apiKey = apiKey
	us.apiSecretKey = apiSecretKey

	viper.Set("ConsumerAPIKey", us.apiKey)
	viper.Set("ConsumerAPISecretKey", us.apiSecretKey)
	return viper.WriteConfig()
}

func (us *UserSecrets) LoadAccessSecrets() error {
	us.accessToken = viper.GetString("AccessToken")
	if len(us.accessToken) == 0 {
		return APIKeyMissingError
	}

	us.accessSecret = viper.GetString("AccessSecret")
	if len(us.accessSecret) == 0 {
		return APIKeyMissingError
	}

	return nil
}

func (us *UserSecrets) SetAccessSecrets(accessToken, accessSecret string) error {
	us.accessToken = accessToken
	us.accessSecret = accessSecret

	viper.Set("AccessToken", us.accessToken)
	viper.Set("AccessSecret", us.accessSecret)
	return viper.WriteConfig()
}

func (us *UserSecrets) GetConsumerKey() string {
	return us.apiKey
}

func (us *UserSecrets) GetConsumerSecret() string {
	return us.apiSecretKey
}

func (us *UserSecrets) GetAccessToken() string {
	return us.accessToken
}

func (us *UserSecrets) GetAccessSecret() string {
	return us.accessSecret
}
