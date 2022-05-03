package rest

import (
	"log"
	"os"

	dotenv "github.com/joho/godotenv"
)

// THIS FUNCTION NEEDS TO BE AMENDED OR REPLACED BECAUSE OF *util/find.go*
func LoadCreds(relativeFileLocation string) {
	if relativeFileLocation == "" {
		relativeFileLocation = "./creds.env"
	}
	err := dotenv.Load(relativeFileLocation)
	if err != nil {
		log.Fatal("Error loading creds.env file")
	}

}

var ReadonlyKey string = os.Getenv("FTX_KEY")
var ReadonlySecret string = os.Getenv("FTX_SECRET")
