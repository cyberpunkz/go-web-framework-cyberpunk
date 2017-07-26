package system

import (
	"os"
	"sync"
	"github.com/cyberpunkz/cyberpunk/logger"
)

type Config struct {
	frontendPath string
}

var instance *Config
var once sync.Once

func GetInstance() *Config {
	once.Do(func() {
		instance = &Config{}
	})

	return instance
}

func (instance *Config) SetFrontendDir(dir string) error {
	appDir, err := os.Getwd()
	instance.frontendPath = appDir + "/" + dir

	if err != nil {
		logger.Error(err)
	}

	return err
}

func (instance *Config) GetFrontendDir() string {
	return instance.frontendPath
}
