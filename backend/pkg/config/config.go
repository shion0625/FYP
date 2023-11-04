package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/caarlos0/env/v6"
)


type Config struct {
	DBHost                string `env:"DB_HOST" envDefault:"hoge_test"`
	DBName                string `env:"DB_NAME" envDefault:"hoge_test"`
	DBUser                string `env:"DB_USER" envDefault:"hoge_test"`
	DBPassword            string `env:"DB_PASSWORD" envDefault:"hoge_test"`
	DBPort                string `env:"DB_PORT" envDefault:"hoge_test"`
	DBSslmode             string `env:"DB_SSLMODE" envDefault:"hoge_test"`
	DBTimezone           string `env:"DB_TIMEZONE" envDefault:"hoge_test"`
	Port                  string `env:"PORT" envDefault:"hoge_test"`
	FrontendUrl           string `env:"FRONTEND_URL" envDefault:"hoge_test"`
	BackendUrl            string `env:"PORT" envDefault:"hoge_test"`
	FrontendDevelopUrl    string `env:"PORT" envDefault:"hoge_test"`
}

var cfg Config

// パッケージ初期化時に一度だけ環境変数をロードし、cfg 変数にセットします
func LoadConfig() (cfg *Config, err error) {

	if err := godotenv.Load(".env"); err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
		return nil, err
	}

    if err := env.Parse(&cfg); err != nil {
		panic(fmt.Sprintf("failed to setup config, error: %s", err))
		return nil, err
	}
	return cfg, nil
}
