package helpers

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func initSignalHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for s := range c {
			log.Printf("Received signal %s, exiting...\n", s)
			os.Exit(0)
		}
	}()
}

func initFlags() {
	flag.Parse()
}

func getEnvOrDefault(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}