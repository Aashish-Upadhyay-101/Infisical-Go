package infisical

import (
	"net/http"
	"time"
)

type SecretBundle struct {
	SecretName    string    `json:"secretName"`
	SecretValue   string    `json:"secretValue"`
	Version       int       `json:"version,omitempty"`
	Workspace     string    `json:"workspace,omitempty"`
	Environment   string    `json:"environment,omitempty"`
	Type          string    `json:"type,omitempty"`
	CreatedAt     string    `json:"createdAt,omitempty"`
	UpdatedAt     string    `json:"updatedAt,omitempty"`
	IsFallback    bool      `json:"isFallBack"`
	LastFetchedAt time.Time `json:"lastFetchAt"`
}

type WorkspaceConfig struct {
	workspaceId  string
	environment  string
	workspaceKey string
}

type ClientConfig struct {
	apiRequest http.Client
	cacheTTL   int32
}

type ServiceTokenCredentials struct {
	serviceTokenKey string
}

type ServiceTokenClientConfig struct {
	ClientConfig
	authMode        string
	credentials     ServiceTokenCredentials
	workspaceConfig WorkspaceConfig
}
