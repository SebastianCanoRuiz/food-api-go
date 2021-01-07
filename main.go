package main

import (
	_ "github.com/joho/godotenv/autoload"
	"go-api/infrastructure"
	"log"
	"os"
)
func main() {
	log.Println("stating API")
	port := os.Getenv("API_PORT")
	infrastructure.Start(port)
}