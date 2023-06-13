package config

import (
	"errors"
	"strings"

	"github.com/geeksheik9/common-functions/pkg/models"
)

// New function creates a new Config object from the given ConfigAccessor and envMap
// When creating new URLs and db URIs in envMap follow pattern host=<url/uri>
func New(ca models.ConfigAccessor, file string) (*models.Config, error) {
	ca.SetConfigFile(file)
	err := ca.ReadInConfig()
	if err != nil {
		return nil, errors.New("Error reading config from file: " + err.Error())
	}

	urls := ca.GetString("URLs")
	var urlArray []string
	if strings.Trim(urls, " ") != "" {
		//url input should be stored in format key=value,key=value,...
		urlArray = strings.Split(urls, ",")
	}
	urlMap := makeMap(urlArray)

	dbs := ca.GetString("DATABASES")
	var dbArray []string
	if strings.Trim(dbs, " ") != "" {
		//url input should be stored in format key=value,key=value,...
		dbArray = strings.Split(dbs, ",")
	}
	dbMap := makeMap(dbArray)

	config := &models.Config{
		Port:        ca.GetString("PORT"),
		LogLevel:    ca.GetString("LOG_LEVEL"),
		URLS:        urlMap,
		Databases:   dbMap,
		SecretsPath: ca.GetString("SECRETS_PATH"),
		CertPath:    ca.GetString("CERT_PATH"),
		DNREmail:    ca.GetString("DO_NOT_REPLY"),
		CourierKey:  ca.GetString("COURIER_KEY"),
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
