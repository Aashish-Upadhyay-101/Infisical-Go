package infisical

import (
	"strings"
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
	if token != "" {
		c := NewClientWithConfig(token, Config{})
		return c
	}
	return nil
}

func NewClientWithConfig(token string, config Config) Client {
	lastDotIdx := strings.LastIndex(token, ".")
	serviceToken := token[0:lastDotIdx]

	config = Config{
		clientConfig: ServiceTokenClientConfig{
			authMode: AUTH_MODE_SERVICE_TOKEN,
			credentials: ServiceTokenCredentials{
				serviceTokenKey: token[lastDotIdx+1:],
			},
			ClientConfig: ClientConfig{
				apiRequest: *CreateApiRequestWithAuthInterceptor(INFISICAL_URL, serviceToken),
				cacheTTL:   config.clientConfig.cacheTTL,
			},
		},
	}

	c := client{
		Config:   config,
		token:    token,
		siteURL:  INFISICAL_URL,
		debug:    false,
		cacheTTL: 300,
	}
	return c
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
