package env

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"os"
)

type Env struct {
	RabbitmqUrl         string
	JwtNotifyPrivateKey string
	JwtKid              string
	JwtIssuer           string
	JwtSubject          string
	JwtAudience         string
}

func LoadEnv(path string) *Env {
	_ = godotenv.Load()

	return &Env{
		RabbitmqUrl:         getEnv("RABBITMQ_URL", "RABBITMQ_URL_LOCALHOST"),
		JwtNotifyPrivateKey: getEnv("JWT_NOTIFY_PRIVATE_KEY", ""),
		JwtKid:              getEnv("JWT_KID", ""),
		JwtIssuer:           getEnv("JWT_ISSUER", ""),
		JwtSubject:          getEnv("JWT_SUBJECT", ""),
		JwtAudience:         getEnv("JWT_AUDIENCE", ""),
	}

}

func getEnv(key, defaultKey string) string {

	if defaultKey != "" {
		key = defaultKey
	}

	val := os.Getenv(key)
	if val == "" {
		log.Error().Msgf("Environment variable %s is not set", key)
		panic("Environment variable not set")
		return ""
	}
	return val
}
