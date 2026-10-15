package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
)

func main() {

	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// collecting data from the web
	c := colly.NewCollector(
		// allow domains
		colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org"),
	)

	fmt.Println("Web Scraper")

	PortString := os.Getenv("Port")
	if PortString == "" {
		log.Fatalf("Port is not available in env")
	}

	fmt.Println("Port: ", PortString)
}
