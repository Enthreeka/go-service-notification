package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type (
	Config struct {
		Postgres                Postgres                `json:"postgres"`
		NotificationHTTTPServer NotificationHTTTPServer `json:"http_server"`
		MailServer              MailServer              `json:"mail_server"`
		ExternalAPI             ExternalAPI             `json:"external_api"`
	}

	Postgres struct {
		URL string `json:"url"`
	}

	NotificationHTTTPServer struct {
		Hostname   string `json:"hostname"`
		Port       string `json:"port"`
		TypeServer string `json:"type_server"`
	}

	MailServer struct {
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
		NotificationHTTTPServer: NotificationHTTTPServer{
			Hostname:   os.Getenv("NOTIFICATION_HTTP_HOSTNAME"),
			Port:       os.Getenv("NOTIFICATION_HTTP_PORT"),
			TypeServer: os.Getenv("NOTIFICATION_HTTP_TYPE_SERVER"),
		},
		MailServer: MailServer{
			Hostname:   os.Getenv("MAIL_HTTP_HOSTNAME"),
			Port:       os.Getenv("MAIL_HTTP_PORT"),
			TypeServer: os.Getenv("MAIL_HTTP_TYPE_SERVER"),
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
