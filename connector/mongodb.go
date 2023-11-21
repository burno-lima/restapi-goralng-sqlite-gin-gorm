package connector

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongodb() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	url := os.Getenv("MONGODB_URL")
	log.Println(url)

	if url == "" {
		log.Fatal("You must set your 'MONGODB_URL' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
}
