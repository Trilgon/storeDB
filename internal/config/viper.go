package config

import "github.com/spf13/viper"

// InitViper - инициализация viper для чтения конфигов
func InitViper() error {
	viper.AddConfigPath("internal/config")
	viper.SetConfigName("configs")
	return viper.ReadInConfig()
}
