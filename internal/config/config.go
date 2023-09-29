package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type (
	Config struct {
		Postgres    Postgres    `json:"postgres"`
		HTTPServer  HTTPServer  `json:"http_server"`
		ExternalAPI ExternalAPI `json:"external_api"`
	}

	Postgres struct {
		URL string `json:"url"`
	}

	HTTPServer struct {
		Hostname   string `json:"hostname"`
		Port       string `json:"port"`
		TypeServer string `json:"type_server"`
	}

	ExternalAPI struct {
		JWT string `json:"jwt"`
	}
)

func New() (*Config, error) {
	err := godotenv.Load("configs/config.env")
	if err != nil {
		return nil, err
	}

	config := &Config{
		Postgres: Postgres{
			URL: os.Getenv("POSTGRES_URL"),
		},
		HTTPServer: HTTPServer{
			Hostname:   os.Getenv("HTTP_HOSTNAME"),
			Port:       os.Getenv("HTTP_PORT"),
			TypeServer: os.Getenv("HTTP_TYPE_SERVER"),
		},
		ExternalAPI: ExternalAPI{
			JWT: os.Getenv("JWT"),
		},
	}

	return config, nil
}

func parseEnvInt(value string) int {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return intValue
}
