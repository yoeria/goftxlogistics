package main

import (
	"log"

	dotenv "github.com/joho/godotenv"
)

// Loads credentials from file into process environment variables
func LoadCreds(root string) (fileLocation string) {
	fileLocation, err := FindFile(root, "creds.env")
	if err != nil {
		return
	}
	err = dotenv.Load(fileLocation)
	if err != nil {
		log.Fatal("Error loading creds.env file")
	}
	return
}
