package util

import "errors"

type UserSecrets struct {
	apiKey       string
	apiSecretKey string
	accessToken  string
	accessSecret string
}

var APIKeyMissingError = errors.New("API Key Missing")

func (us *UserSecrets) LoadSecrets() error {
	// load API key from storage
	if len(us.apiKey) == 0 {
		return APIKeyMissingError
	}

	// load API secret key from storage
	if len(us.apiSecretKey) == 0 {
		return APIKeyMissingError
	}

	// load access token from storage
	if len(us.accessToken) == 0 {
		return APIKeyMissingError
	}

	return nil
}

func (us *UserSecrets) SetSecrets(apiKey, apiSecretKey string) error {
	us.apiKey = apiKey
	us.apiSecretKey = apiSecretKey

	// save them to persistant storage

	return nil
}

func (us *UserSecrets) SetAccessSecrets(accessToken, accessSecret string) error {
	us.accessToken = accessToken
	us.accessSecret = accessSecret

	// save them to persistant storage

	return nil
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
