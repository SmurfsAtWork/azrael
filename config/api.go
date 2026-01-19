package config

import "net/url"

func ApiAddress() string {
	return configValues.Api.Address
}

func SetApiAddress(address string) error {
	_, err := url.ParseRequestURI(address)
	if err != nil {
		return err
	}
	configValues.Api.Address = address
	return saveConfigFile()
}

func ResetApiAddress() error {
	return SetApiAddress(defaultConfig.Api.Address)
}
