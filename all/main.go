package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func init() {
	log.Println("init in main")
}

func main() {
	// load configuration
	configFile := flag.String("config", "./all.yml", "Configuration file path.")

	var cfg Config
	if err := cfg.FromFile(*configFile); err != nil {
		fmt.Fprintf(os.Stderr, "jackal: %v\n", err)
		return
	}

	log.Printf("%v", cfg)
}
