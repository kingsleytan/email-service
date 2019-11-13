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
)

func getEnv(i string) string {
	l, isExist := os.LookupEnv(i)
	if !isExist {
		log.Printf("%s cannot find the environment", i)
	}

	return l
}
