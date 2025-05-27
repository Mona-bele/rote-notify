package utils

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type RMQConfig struct {
	RabbitmqUrl               string
	RabbitmqMaxReconnectTimes int
}

type Env struct {
	POSTGRES_HOST        string `mapstructure:"DB_HOST"`
	POSTGRES_PORT        string `mapstructure:"DB_PORT"`
	POSTGRES_USERNAME    string `mapstructure:"DB_USERNAME"`
	POSTGRES_PASSWORD    string `mapstructure:"DB_PASSWORD"`
	POSTGRES_DATABASE    string `mapstructure:"DB_NAME"`
	POSTGRES_SSL         string `mapstructure:"SSL_MODE"`
	DATABASE_URL_PRIMARY string `mapstructure:"DATABASE_URL_PRIMARY"`
	LOG_DEBUG            string `mapstructure:"LOG_DEBUG"`
	DATABASE_URL_MONGO   string `mapstructure:"DATABASE_URL_MONGO"`

	PORT_EVENT string `mapstructure:"PORT_EVENT"`

	QUEUE_NAME_UPDATE_USER_EVENT    string `mapstructure:"QUEUE_NAME_UPDATE_USER_EVENT"`
	EXCHANGE_NAME_UPDATE_USER_EVENT string `mapstructure:"EXCHANGE_NAME_UPDATE_USER_EVENT"`

	DATABASE_URL_WALLET string `mapstructure:"DATABASE_URL_WALLET"`

	RMQConfig *RMQConfig
}

func LoadEnv(path string) *Env {
	_ = godotenv.Load()

	viper.SetConfigFile(path)
	viper.AutomaticEnv()
	_ = viper.ReadInConfig()

	databaseUrlPrimary := viper.GetString("DATABASE_URL_PRIMARY")
	DatabaseUrlWallet := viper.GetString("DATABASE_URL_WALLET")
	MongoDatabaseDns := viper.GetString("MONGO_URL")
	port := viper.GetString("PORT_EVENT")
	PostgresHost := viper.GetString("POSTGRES_HOST")
	PostgresPort := viper.GetString("POSTGRES_PORT")
	PostgresUser := viper.GetString("POSTGRES_USERNAME")
	PostgresPass := viper.GetString("POSTGRES_PASSWORD")
	PostgresDbname := viper.GetString("POSTGRES_DATABASE")
	PostgresSslMode := viper.GetString("SSL_MODE")
	LogDebug := viper.GetString("LOG_DEBUG")
	rmq_url := viper.GetString("RABBITMQ_URL")
	QueuNameUpdateUserEvent := viper.GetString("QUEUE_NAME_UPDATE_USER_EVENT")
	ExchangeNameUpdateUserEvent := viper.GetString("EXCHANGE_NAME_UPDATE_USER_EVENT")

	env := &Env{
		PORT_EVENT:           port,
		DATABASE_URL_PRIMARY: databaseUrlPrimary,
		POSTGRES_HOST:        PostgresHost,
		POSTGRES_PORT:        PostgresPort,
		POSTGRES_USERNAME:    PostgresUser,
		POSTGRES_PASSWORD:    PostgresPass,
		POSTGRES_DATABASE:    PostgresDbname,
		POSTGRES_SSL:         PostgresSslMode,
		LOG_DEBUG:            LogDebug,
		DATABASE_URL_WALLET:  DatabaseUrlWallet,
		DATABASE_URL_MONGO:   MongoDatabaseDns,
		RMQConfig: &RMQConfig{
			RabbitmqUrl: rmq_url,
		},
		QUEUE_NAME_UPDATE_USER_EVENT:    QueuNameUpdateUserEvent,
		EXCHANGE_NAME_UPDATE_USER_EVENT: ExchangeNameUpdateUserEvent,
	}

	return env
}
