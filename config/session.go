package config

func SessionToken() string {
	return configValues.Session.Token
}

func SetSessionToken(token string) error {
	configValues.Session.Token = token
	return saveConfigFile()
}
