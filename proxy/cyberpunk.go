package proxy

import (
	"github.com/cyberpunkz/cyberpunk/config/system"
	"github.com/cyberpunkz/cyberpunk/logger"
)

func GetSystemConfig() *system.Config {
	return system.GetInstance()
}

func Log() logger.Logger {
	return GetSystemConfig().GetCurrentLogger()
}

