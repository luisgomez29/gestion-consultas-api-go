package config

import "github.com/spf13/viper"

// Config variables de entorno
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

// ServerConfig variables de entorno de servidor
type ServerConfig struct {
	Port string
}

// DatabaseConfig variables de entorno de la base de datos
type DatabaseConfig struct {
	Host     string
	Name     string
	User     string
	Password string
	Port     string
}

// Load lee las configuraciones del archivo de configuración dentro de la ruta si existe.
func Load(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
