package models

import "crypto/x509"

// Config is the main configuration object
// swagger:model Config
type Config struct {
	Port        string            `json:"port"`
	LogLevel    string            `json:"logLevel"`
	URLS        map[string]string `json:"urls"`
	Databases   map[string]string `json:"databases"`
	SecretsPath string            `json:"secretsPath"`
	CertPath    string            `json:"certPath"`
	CertPool    *x509.CertPool    `json:"certPool"`
	DNREmail    string            `json:"doNotReplyEmail"`
}

// ConfigAccessor is the basic interface that configuration objects must implement
// swagger:model ConfigAccessor
type ConfigAccessor interface {
	BindEnv(input ...string) error
	IsSet(key string) bool
	GetString(key string) string
}
