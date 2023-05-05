package infisical

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
	token   string
	siteURL string
	debug   bool
}
