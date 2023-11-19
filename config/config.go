package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/phamdinhha/go-chat-server/pkg/constants"
	"github.com/phamdinhha/go-chat-server/pkg/logger"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "BankAccount microservice config path")
}

type Config struct {
	ServiceName string           `mapstructure:"serviceName"`
	Logger      logger.LogConfig `mapstructure:"logger"`
	JWTConfig   JWTConfig        `mapstructure:"jwt"`
}

type JWTConfig struct {
	JWTTTL        int64  `mapstructure:"jwtTTL"`
	JWTSecret     string `mapstructure:"jwtSecret"`
	SigningMethod string `mapstructure:"signingMethod"`
}

func InitConfig() (*Config, error) {
	if configPath == "" {
		configPathFromEnv := os.Getenv(constants.ConfigPath)
		if configPathFromEnv != "" {
			configPath = configPathFromEnv
		} else {
			getwd, err := os.Getwd()
			if err != nil {
				return nil, errors.Wrap(err, "os.Getwd")
			}
			configPath = fmt.Sprintf("%s/config/config.yaml", getwd)
		}
	}

	cfg := &Config{}

	viper.SetConfigType(constants.Yaml)
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "viper.ReadInConfig")
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, errors.Wrap(err, "viper.Unmarshal")
	}

	return cfg, nil
}
