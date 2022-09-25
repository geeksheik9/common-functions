package config

import (
	"errors"
	"strings"

	"github.com/geeksheik9/common-functions/pkg/models"
)

// New function creates a new Config object from the given ConfigAccessor and envMap
// When creating new URLs and db URIs in envMap follow pattern host=<url/uri>
func New(ca models.ConfigAccessor, envMap map[string]string) (*models.Config, error) {
	err := loadEnvVars(ca, envMap)
	if err != nil {
		return nil, err
	}

	var urlArray []string
	if _, ok := envMap[URLs]; ok {
		//url input should be stored in format key=value,key=value,...
		urlArray = strings.Split(envMap[URLs], ",")
	}
	urlMap := makeMap(urlArray)

	var dbArray []string
	if _, ok := envMap[databases]; ok {
		//db input should be stored in format key=value,key=value,...
		dbArray = strings.Split(envMap[databases], ",")
	}
	dbMap := makeMap(dbArray)

	config := &models.Config{
		Port:        envMap[port],
		LogLevel:    envMap[logLevel],
		URLS:        urlMap,
		Databases:   dbMap,
		SecretsPath: envMap[secretsPath],
		CertPath:    envMap[certPath],
		DNREmail: envMap[doNotReply],
	}

	return config, nil
}

// makeMap function creates a map from the given array of strings, inteded for use with URLs and databases
// When creating new URLs and db URIs in envMap follow pattern host=<url/uri>
func makeMap(array []string) map[string]string {
	m := make(map[string]string)
	for _, v := range array {
		pair := strings.Split(v, "=")
		m[pair[0]] = pair[1]
	}
	return m
}

// loadEnvVars function loads the env vars from the given ConfigAccessor and envMap
// helper function for New
func loadEnvVars(ca models.ConfigAccessor, envMap map[string]string) error {
	for key := range envMap {
		err := ca.BindEnv(key)
		if err != nil {
			return errors.New("failed to bind env var " + key)
		}

		if ca.IsSet(key) {
			envMap[key] = ca.GetString(key)
		} else {
			return errors.New("env var " + key + " is not set")
		}
	}
	return nil
}
