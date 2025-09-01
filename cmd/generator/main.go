package main

import (
	"log"
	"os"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
)

const GENERATED = "generated.go"

func main() {
	config, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	err = api.Generate(config)
	if err != nil {
		log.Fatalf("failed to generate: %v", err)
	}

	// generaeted omits root models, so we need to remove the file
	os.Remove(GENERATED)

	log.Println("models generated")
}
