package env

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"os"
)

type Env struct {
	RabbitmqUrl string
}

func LoadEnv(path string) *Env {
	_ = godotenv.Load(path)

	return &Env{
		RabbitmqUrl: getEnv("RABBITMQ_URL"),
	}

}

func getEnv(key string) string {

	val := os.Getenv(key)
	if val == "" {
		log.Error().Msgf("Environment variable %s is not set", key)
		panic("Environment variable not set")
	}
	return val
}
