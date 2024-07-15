package config

import (
	"errors"
	"reflect"

	"github.com/geeksheik9/common-functions/pkg/models"
	"github.com/geeksheik9/common-functions/pkg/utils"
)

// New function creates a new Config object from the given ConfigAccessor and envMap
// When creating new URLs and db URIs in envMap follow pattern host=<url/uri>
func New(ca models.ConfigAccessor, file string, configurationModel interface{}, configurationValues []string) (interface{}, error) {
	ca.SetConfigFile(file)
	err := ca.ReadInConfig()
	if err != nil {
		return nil, errors.New("Error reading config from file: " + err.Error())
	}

	config := &models.Config{
		Port:     ca.GetString("PORT"),
		LogLevel: ca.GetString("LOG_LEVEL"),
	}

	c := models.Config{}

	var values []string
	for _, value := range configurationValues {
		values = append(values, ca.GetString(value))
	}

	if reflect.TypeOf(c) != configurationModel {
		fields := utils.GetFields(configurationModel)
		config.AddFields(fields, values)
	}

	return config, nil
}
