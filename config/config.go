package config

import (
	"os"
	"os/user"

	"github.com/SmurfsAtWork/azrael/log"
	"gopkg.in/yaml.v3"
)

var (
	configFilePath   string
	configFileHeader = `# why are you reading this comment?
# missing something?
# it's a yaml file, don't bother with comments!
# stop
# ok if you insist...
#
# This is the configuration of azrael, don't play with it unless you know what you're doing :)

`
	configValues  config
	defaultConfig = config{
		Api: struct{ Address string }{
			Address: "https://papa.smurfsatwork.org",
		},
		Session: struct{ Token string }{
			Token: "",
		},
	}
)

func init() {
	u, _ := user.Current()
	configFilePath = u.HomeDir + "/.azrael.yaml"

	err := loadConfigFile()
	if err != nil {
		configValues = defaultConfig
		err = saveConfigFile()
		if err != nil {
			log.Fatalln(err)
		}
	}
}

type config struct {
	Api struct {
		Address string
	}
	Session struct {
		Token string
	}
}

func loadConfigFile() error {
	configFile, err := os.Open(configFilePath)
	if os.IsNotExist(err) {
		log.Infoln("Config file was not found at ~/.azrael.yaml creating new file with default configs!")
		configFile, err = createDefaultConfigFile()
		if err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}
	err = yaml.NewDecoder(configFile).Decode(&configValues)
	defer configFile.Close()
	if err != nil {
		return err
	}
	return nil
}

func saveConfigFile() error {
	configFile, err := os.OpenFile(configFilePath, os.O_RDWR|os.O_APPEND, os.ModeAppend)
	if os.IsNotExist(err) {
		log.Infoln("Config file was not found at ~/.azrael.yaml creating new file with default configs!")
		configFile, err = createDefaultConfigFile()
		if err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}

	_ = configFile.Truncate(0)
	_, err = configFile.WriteString(configFileHeader)
	if err != nil {
		return err
	}

	err = yaml.NewEncoder(configFile).Encode(configValues)
	defer configFile.Close()
	if err != nil {
		return err
	}
	return nil
}

func createDefaultConfigFile() (*os.File, error) {
	configFile, err := os.Create(configFilePath)
	if err != nil {
		return nil, err
	}

	err = yaml.NewEncoder(configFile).Encode(defaultConfig)
	if err != nil {
		return nil, err
	}

	return configFile, nil
}
