package models

import (
	"fmt"
	"reflect"
)

// Config is the main configuration object
// swagger:model Config
type Config struct {
	Port     string `json:"port"`
	LogLevel string `json:"logLevel"`
	// URLS        map[string]string `json:"urls"`
	// Databases   map[string]string `json:"databases"`
	// SecretsPath string            `json:"secretsPath"`
	// CertPath    string            `json:"certPath"`
	// CertPool    *x509.CertPool    `json:"certPool"`
}

func (c *Config) AddFields(fields []string, values []string) interface{} {
	defaultFields := reflect.ValueOf(*c)
	for i := 0; i < defaultFields.NumField(); i++ {

		//removes Config field from fields and values
		if fields[i] == "Config" {
			fields[i] = fields[len(fields)-1]
			fields[len(fields)-1] = ""
			fields = fields[:len(fields)-1]

			values[i] = values[len(values)-1]
			values[len(values)-1] = ""
			values = values[:len(values)-1]
		}
		fields = append(fields, defaultFields.Type().Field(i).Name)
		fmt.Println("fields", fields)
		fmt.Println("values", values)
	}

	conf := []reflect.StructField{}
	for i, field := range fields {
		x := reflect.StructField{
			Name: field,
			Type: reflect.TypeOf(values[i]),
		}
		conf = append(conf, x)
	}

	config := reflect.StructOf(conf)

	//fmt.Println(conf)

	for i, field := range fields {
		f, err := config.FieldByName(field)
		if err {
			return nil
		}

		t := reflect.ValueOf(&f)

		fmt.Println("1", t)
		if t.CanSet() {
			t.Set(reflect.ValueOf(values[i]))
		}
		fmt.Println("2", t)
	}

	return config
}

type Conf interface {
	Config
}

// ConfigAccessor is the basic interface that configuration objects must implement
// swagger:model ConfigAccessor
type ConfigAccessor interface {
	BindEnv(input ...string) error
	IsSet(key string) bool
	GetString(key string) string
	SetConfigFile(file string)
	ReadInConfig() error
}
