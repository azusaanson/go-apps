package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBUser       string `mapstructure:"DB_USER"`
	DBPassword   string `mapstructure:"DB_PASSWORD"`
	DBHost       string `mapstructure:"DB_HOST"`
	DBPort       string `mapstructure:"DB_PORT"`
	DBName       string `mapstructure:"DB_NAME"`
	MigrationURL string `mapstructure:"MIGRATION_URL"`

	GRPCServer string `mapstructure:"GRPC_SERVER"`

	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
