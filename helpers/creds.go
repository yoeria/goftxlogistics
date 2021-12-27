package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main(){
	LoadCreds()
}

// Give "key" or "secret" as argument for desired return value
func LoadCreds()  {

	err := godotenv.Load("./creds.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}
var ReadonlyKey string = os.Getenv("FTX_KEY")
var ReadonlySecret string = os.Getenv("FTX_SECRET")
