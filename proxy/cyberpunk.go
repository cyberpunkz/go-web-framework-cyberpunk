package proxy

import (
	"github.com/cyberpunkz/cyberpunk/config/system"
	"log"
)

func GetSystemConfig() *system.Config {
	return system.GetInstance()
}

func Error() *log.Logger {
	return GetSystemConfig().GetCurrentLogger().Error()
}

func Debug() *log.Logger {
	return GetSystemConfig().GetCurrentLogger().Debug()
}
