package config

import (
	"os"
	"time"
)

type Config struct {
	APP         string
	Environment string
	LogLevel    string
	RPCPort     string
	RabbitMQURL string
	Context     struct {
		Timeout string
	}
	Token struct {
		Secret     string
		AccessTTL  time.Duration
		RefreshTTL time.Duration
	}

	DB struct {
		Host                string
		Port                string
		Name                string
		User                string
		Password            string
		CollectionName      string
		AlarmCollectionName string
		SpeakerCollection   string
	}
}

func Token() string {
	c := Config{}
	c.Token.Secret = getEnv("TOKEN_SECRET", "token_secret")
	return c.Token.Secret
}

func New() *Config {
	var config Config

	config.APP = getEnv("APP", "app")
	config.Environment = getEnv("ENVIRONMENT", "develop")
	config.LogLevel = getEnv("LOG_LEVEL", "local")
	config.RPCPort = getEnv("RPC_PORT", "device_service:9001")
	config.Context.Timeout = getEnv("CONTEXT_TIMEOUT", "30s")

	config.DB.Host = getEnv("MONGO_HOST", "mongo")
	config.DB.Port = getEnv("MONGO_PORT", ":27017")
	config.DB.User = getEnv("MONGO_USER", "")
	config.DB.Password = getEnv("MONGO_PASSWORD", "")
	config.DB.CollectionName = getEnv("MONGO_COLLECTION", "tv")
	config.DB.AlarmCollectionName = getEnv("MONGO_ALARM_COLLECTION", "alarm")
	config.DB.SpeakerCollection = getEnv("MONGO_SPEAKER_COLLECTION", "speaker")
	config.DB.Name = getEnv("MONGO_DATABASE", "devices")

	config.Token.Secret = getEnv("TOKEN_SECRET", "D1YORTOP4EEK")
	accessTTl, err := time.ParseDuration(getEnv("TOKEN_ACCESS_TTL", "1h"))
	if err != nil {
		return nil
	}
	refreshTTL, err := time.ParseDuration(getEnv("TOKEN_REFRESH_TTL", "24h"))
	if err != nil {
		return nil
	}
	config.Token.AccessTTL = accessTTl
	config.Token.RefreshTTL = refreshTTL

	config.RabbitMQURL = getEnv("RabbitMQ_URL", "amqp://guest:guest@rabbitmq:5672/")
	return &config
}

func getEnv(key string, defaultVaule string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultVaule
}
