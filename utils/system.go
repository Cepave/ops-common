package utils

import (
	"log"
	"os"
)

func Hostname(configHostname string) (string, error) {
	if configHostname != "" {
		return configHostname, nil
	}

	hostname, err := os.Hostname()
	if err != nil {
		log.Println("ERROR: os.Hostname() fail", err)
	}

	return hostname, err
}
