package main

import (
	log "github.com/Sirupsen/logrus"
)

func run() error {
	return nil
}

func main() {
	log.Info("Starting script")
	if err := run(); err != nil {
		log.Panic("...")
	}

}
