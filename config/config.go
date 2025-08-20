package config

import (
	"fmt"
	"log"
	"os"
	"supply_chain_platform/logger"
	"sync"

	"github.com/joho/godotenv"
)

type Configuration struct {
	Environment      string
	AdminServicePort string
	AuthServicePort  string
	NotificationPort string
	PaymentPort      string
	ShipmentPort     string
	ThrottleTTL      string
	ThrottleLimit    string
	DBDialect        string
	DBUsername       string
	DBPassword       string
	DBHost           string
	DBPort           string
	RedisURL         string
	KafkaBrokers     string
	StripeSecretKey  string
	JWTSecret        string
	GoogleMapsAPIKey string
	WeatherAPIKey    string
	GinMode          string
}

var (
	AppConfig Configuration
	initOnce  sync.Once
	envLoaded bool
)

func init() {
	initOnce.Do(func() {
		env := os.Getenv("GO_ENV")
		if env == "" {
			env = "dev" // default
		}

		envFile := ".env." + env
		fmt.Printf("Loading environment from: %s\n", envFile)

		err := godotenv.Load(envFile)
		if err != nil {
			log.Fatalf("Error loading environment file (%s): %v", envFile, err)
		}

		AppConfig = Configuration{
			Environment:      os.Getenv("ENVIRONMENT"),
			AdminServicePort: os.Getenv("ADMIN_SERVICE_PORT"),
			AuthServicePort:  os.Getenv("AUTH_SERVICE_PORT"),
			NotificationPort: os.Getenv("NOTIFICATION_SERVICE_PORT"),
			PaymentPort:      os.Getenv("PAYMENT_SERVICE_PORT"),
			ShipmentPort:     os.Getenv("SHIPMENT_SERVICE_PORT"),
			ThrottleTTL:      os.Getenv("THROTTLE_TTL"),
			ThrottleLimit:    os.Getenv("THROTTLE_LIMIT"),
			DBDialect:        os.Getenv("DB_DIALECT"),
			DBUsername:       os.Getenv("DB_USERNAME"),
			DBPassword:       os.Getenv("DB_PASSWORD"),
			DBHost:           os.Getenv("DB_HOST"),
			DBPort:           os.Getenv("DB_PORT"),
			RedisURL:         os.Getenv("REDIS_URL"),
			KafkaBrokers:     os.Getenv("KAFKA_BROKERS"),
			StripeSecretKey:  os.Getenv("STRIPE_SECRET_KEY"),
			JWTSecret:        os.Getenv("JWT_SECRET"),
			GoogleMapsAPIKey: os.Getenv("GOOGLE_MAPS_API_KEY"),
			WeatherAPIKey:    os.Getenv("WEATHER_API_KEY"),
			GinMode:          os.Getenv("GIN_MODE"),
		}

		log.Printf("Environment: %s", AppConfig.Environment)
		envLoaded = true
	})

	if !envLoaded {
		logger.GetLogger("config").Fatalf("Environment not loaded")
	}
}
