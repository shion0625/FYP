package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	DBHost             string `env:"DB_HOST"              envDefault:"hoge_test"`
	DBName             string `env:"DB_NAME"              envDefault:"hoge_test"`
	DBUser             string `env:"DB_USER"              envDefault:"hoge_test"`
	DBPassword         string `env:"DB_PASSWORD"          envDefault:"hoge_test"`
	DBPort             string `env:"DB_PORT"              envDefault:"hoge_test"`
	DBSslmode          string `env:"DB_SSLMODE"           envDefault:"hoge_test"`
	DBTimezone         string `env:"DB_TIMEZONE"          envDefault:"hoge_test"`
	Port               string `env:"PORT"                 envDefault:"hoge_test"`
	FrontendUrl        string `env:"FRONTEND_URL"         envDefault:"hoge_test"`
	BackendUrl         string `env:"BackendUrl"           envDefault:"hoge_test"`
	FrontendDevelopUrl string `env:"FrontendDevelopUrl"   envDefault:"hoge_test"`
	AdminAuthKey       string `env:"ADMIN_AUTH_KEY"       envDefault:"hoge_test"`
	UserAuthKey        string `env:"USER_AUTH_KEY"        envDefault:"hoge_test"`
	GcpBucketName      string `env:"GCP_BUCKET_NAME"      envDefault:"hoge_test"`
	CredentialsFile    string `env:"GCP_CREDENTIALS_File" envDefault:"hoge_test"`
	GcpServiceAccount  string `env:"GCP_SERVICE_ACCOUNT" envDefault:"hoge_test"`
}

func LoadConfig() (cfg *Config, err error) {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Printf("Failed to load: %v", err)

		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	cfg = new(Config)
	if err := env.Parse(cfg); err != nil {
		fmt.Printf("failed to setup config, error: %s", err)

		return nil, fmt.Errorf("failed to setup config, error: %w", err)
	}

	return cfg, nil
}
