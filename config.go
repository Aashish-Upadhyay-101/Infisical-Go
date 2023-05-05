package infisical

type Config struct {
	cache        map[string]SecretBundle
	clientConfig ServiceTokenClientConfig
	debug        bool
}

const INFISICAL_URL = "https://app.infisical.com"
