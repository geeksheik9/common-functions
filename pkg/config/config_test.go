package config

import (
	"errors"
	"testing"

	"github.com/geeksheik9/common-functions/pkg/mocks"
)

func Test_MakeMap(t *testing.T) {
	array := []string{"key1=value1", "key2=value2"}
	expected := map[string]string{"key1": "value1", "key2": "value2"}
	actual := makeMap(array)
	if len(actual) != len(expected) {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func Test_New_NoError(t *testing.T) {
	ca := &mocks.ConfigAccessor{}

	envMap := make(map[string]string)
	envMap[port] = "8080"
	envMap[logLevel] = "debug"
	envMap[URLs] = "host1=https://localhost:8080,host2=https://localhost:8081"
	envMap[databases] = "host1=mongoCollection1,host2=SQLTable1"
	envMap[secretsPath] = "/tmp/secrets"
	envMap[certPath] = "/tmp/cert.pem"

	for envKey := range envMap {
		ca.On("BindEnv", envKey).Return(nil)
		ca.On("IsSet", envKey).Return(true)
		ca.On("GetString", envKey).Return(envMap[envKey])
	}

	c, _ := New(ca, envMap)

	if c.Port != "8080" || c.LogLevel != "debug" || c.URLS["host1"] != "https://localhost:8080" || c.Databases["host2"] != "SQLTable1" || c.SecretsPath != "/tmp/secrets" || c.CertPath != "/tmp/cert.pem" {
		t.Errorf("expected: %v\nreceived: %v", envMap, c)
	}
}

func Test_New_BindError(t *testing.T) {
	ca := &mocks.ConfigAccessor{}

	envMap := make(map[string]string)
	envMap[port] = "8080"
	envMap[logLevel] = "debug"
	envMap[URLs] = "host1=https://localhost:8080,host2=https://localhost:8081"
	envMap[databases] = "host1=mongoCollection1,host2=SQLTable1"
	envMap[secretsPath] = "/tmp/secrets"
	envMap[certPath] = "/tmp/cert.pem"

	for envKey := range envMap {
		if envKey == port {
			ca.On("BindEnv", envKey).Return(errors.New("failed to bind env var " + envKey))
		}
		ca.On("BindEnv", envKey).Return(nil)
		ca.On("IsSet", envKey).Return(true)
		ca.On("GetString", envKey).Return(envMap[envKey])
	}

	c, err := New(ca, envMap)
	if err == nil {
		t.Errorf("expected error, received: %v", err)
	}
	if c != nil {
		t.Errorf("expected nil, received: %v", c)
	}
}

func Test_New_SetError(t *testing.T) {
	ca := &mocks.ConfigAccessor{}

	envMap := make(map[string]string)
	envMap[port] = "8080"
	envMap[logLevel] = "debug"
	envMap[URLs] = "host1=https://localhost:8080,host2=https://localhost:8081"
	envMap[databases] = "host1=mongoCollection1,host2=SQLTable1"
	envMap[secretsPath] = "/tmp/secrets"
	envMap[certPath] = "/tmp/cert.pem"

	for envKey := range envMap {
		if envKey == port {
			ca.On("BindEnv", envKey).Return(nil)
			ca.On("IsSet", envKey).Return(false)
		}
		ca.On("BindEnv", envKey).Return(nil)
		ca.On("IsSet", envKey).Return(true)
		ca.On("GetString", envKey).Return(envMap[envKey])
	}

	c, err := New(ca, envMap)
	if err == nil {
		t.Errorf("expected error, received: %v", err)
	}
	if c != nil {
		t.Errorf("expected nil, received: %v", c)
	}
}
