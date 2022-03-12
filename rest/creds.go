package rest

import (
	"log"
	"os"

	dotenv "github.com/joho/godotenv"
)

// Give "key" or "secret" as argument for desired return value
func LoadCreds() {

	err := dotenv.Load("./creds.env")
	if err != nil {
		log.Fatal("Error loading creds.env file")
	}

}

var ReadonlyKey string = os.Getenv("FTX_KEY")
var ReadonlySecret string = os.Getenv("FTX_SECRET")
