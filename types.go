package infisical

import "time"

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
