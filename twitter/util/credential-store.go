package util

import "errors"

type UserSecrets struct {
	apiKey       string
	apiSecretKey string
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

	return nil
}

func (us *UserSecrets) SetSecrets(apiKey, apiSecretKey string) error {
	us.apiKey = apiKey
	us.apiSecretKey = apiSecretKey

	// save them to persistant storage

	return nil
}
