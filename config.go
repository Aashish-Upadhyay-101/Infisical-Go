package infisical

const (
	INFISICAL_URL           = "https://app.infisical.com"
	AUTH_MODE_SERVICE_TOKEN = "serviceToken"
)

type Config struct {
	cache        map[string]SecretBundle
	clientConfig ServiceTokenClientConfig
	debug        bool
}
