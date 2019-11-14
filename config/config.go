package config

import (
	"log"
	"os"
)

// configuration
var (
	Env           = getEnv("ENV")
	MailgunDomain = getEnv("DOMAIN")
	MailgunKey    = getEnv("MAILGUN_KEY")
	DBUserName    = getEnv("DB_USERNAME")
	DBPassword    = getEnv("DB_PASSWORD")
	DBName        = getEnv("DB_DATABASE")
	DBHost        = getEnv("DB_HOST")
	DBPort        = getEnv("DB_PORT")
	DBUnixSocket  = ""
)

func getEnv(i string) string {
	l, isExist := os.LookupEnv(i)
	if !isExist {
		log.Printf("%s cannot find the environment", i)
	}

	return l
}
