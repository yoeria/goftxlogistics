package main

import (
	"log"
	"os"

	dotenv "github.com/joho/godotenv"
	"github.com/yoeria/goftxlogistics/util"
)

// Loads credentials from file into process environment variables \n
//
func LoadCreds(root string) (fileLocation string) {
	fileLocation, err := util.FindFile(root, "creds.env")
	if err != nil {
		return fileLocation
	}
	err = dotenv.Load(fileLocation)
	if err != nil {
		log.Fatal("Error loading creds.env file")
	}
	return
}

var ReadOnlyKey string = os.Getenv("FTX_KEY")
var ReadOnlySecret string = os.Getenv("FTX_SECRET")
