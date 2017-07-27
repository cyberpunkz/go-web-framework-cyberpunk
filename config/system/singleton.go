package system

import (
	"os"
	"sync"
	"github.com/cyberpunkz/cyberpunk/logger"
)

type Config struct {
	logger logger.Logger
	logConfig logger.LogConfig
	frontendPath string
}

var instance *Config
var once sync.Once

func GetInstance() *Config {
	once.Do(func() {
		instance = &Config{}
		instance.DisableLogWriting()
	})

	return instance
}

func (instance *Config) InitLogger(path string) {
	logConfig := logger.LogConfig{Path:path}
	instance.logConfig = logConfig
}

func (instance *Config) EnableLogWriting() {
	instance.logger = logger.GetFile()
	instance.logger.AddLogConfig(instance.logConfig)
}

func (instance *Config) DisableLogWriting() {
	instance.logger = logger.GetConsole()
	instance.logger.AddLogConfig(instance.logConfig)
}

func (instance *Config) GetCurrentLogger() logger.Logger {
	return instance.logger
}

func (instance *Config) SetFrontendDir(dir string) error {
	appDir, err := os.Getwd()
	instance.frontendPath = appDir + "/" + dir

	if err != nil {
		instance.logger.Error().Println(err)
	}

	return err
}

func (instance *Config) GetFrontendDir() string {
	return instance.frontendPath
}
