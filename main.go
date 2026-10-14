package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Println("Web Scraper")

	PortString := os.Getenv("Port")
	if PortString == "" {
		log.Fatal("Port is not available in env")
	}

	fmt.Println("Port: ", PortString)
}
