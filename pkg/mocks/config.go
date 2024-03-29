package mocks

import mock "github.com/stretchr/testify/mock"

type ConfigAccessor struct {
	mock.Mock
}

func (_m *ConfigAccessor) BindEnv(input ...string) error {
	_va := make([]interface{}, len(input))
	for _i := range input {
		_va[_i] = input[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...string) error); ok {
		r0 = rf(input...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *ConfigAccessor) GetString(key string) string {
	ret := _m.Called(key)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

func (_m *ConfigAccessor) IsSet(key string) bool {
	ret := _m.Called(key)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
