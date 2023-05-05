package infisical

import (
	"errors"
)

// this interface is the main API exposed by Infisical Go SDK
type Client interface {
	// method for getting all the secrets
	getAllSecrets() ([]SecretBundle, error)

	// method for getting a secret
	getSecret() (SecretBundle, error)

	// method for creating a secret
	createSecret() (SecretBundle, error)

	// method for updating a secret
	updateSecret() (SecretBundle, error)

	// method for deleting a secret
	deleteSecret() (SecretBundle, error)
}

type client struct {
	Config
	token    string
	siteURL  string
	debug    bool
	cacheTTL int32
}

func InfisicalClient(token string) Client {
	c, err := NewClientWithConfig(token, Config{})
	if err != nil {

	}

	return c
}

func NewClientWithConfig(token string, config Config) (Client, error) {
	if token != "" {
		c := client{
			Config:   config,
			token:    token,
			siteURL:  INFISICAL_URL,
			debug:    false,
			cacheTTL: 300,
		}
		return c, nil
	}
	return nil, errors.New("Token can't be empty")
}

func (c client) getAllSecrets() ([]SecretBundle, error) {
	return []SecretBundle{}, nil
}

func (c client) createSecret() (SecretBundle, error) {
	return SecretBundle{}, nil
}

func (c client) getSecret() (SecretBundle, error) {
	return SecretBundle{}, nil
}

func (c client) updateSecret() (SecretBundle, error) {
	return SecretBundle{}, nil
}

func (c client) deleteSecret() (SecretBundle, error) {
	return SecretBundle{}, nil
}
