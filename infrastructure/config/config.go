package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"database"`
}

// application specific config
type AppConfig struct {
	Port         int `mapstructure:"port"`
	ReadTimeout  int `mapstructure:"read_timeout"`
	WriteTimeout int `mapstructure:"write_timeout"`
	IdleTimeout  int `mapstructure:"idle_timeout"`
}

// database specific config
type DatabaseConfig struct {
	Host        string        `mapstructure:"host"`
	Port        int           `mapstructure:"port"`
	Name        string        `mapstructure:"name"`
	Username    string        `mapstructure:"username"`
	Password    string        `mapstructure:"password"`
	SslMode     string        `mapstructure:"ssl_mode"`
	MaxLifeTime time.Duration `mapstructure:"max_life_time"`
	MaxIdleConn int           `mapstructure:"max_idle_conn"`
	MaxOpenConn int           `mapstructure:"max_open_conn"`
	Debug       bool          `mapstructure:"debug"`
}

// c is the configuration instance
var c Config

// Get returns all configurations
func Get() Config {
	return c
}

// Load config from config file
func Load() error {
	viper.SetConfigName("config.local")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./infrastructure/config/")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read config file: %s \n", err))
	}

	if err := viper.Unmarshal(&c); err != nil {
		return fmt.Errorf("failed to unmarshal config: %v", err)
	}

	return nil
}
