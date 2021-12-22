package main

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Give "key" or "secret" as argument for desired return value
func Creds(i string) {

	err := godotenv.Load("./creds.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	readonlyKey := os.Getenv("FTX_KEY")
	readonlySecret := os.Getenv("FTX_SECRET")

	if i == "" {
		errors.New("Please specify whether to load the key or secret")
	}

	if i == "key"{
		return readonlyKey
	}
	if i == "secret"{
		return readonlySecret
	}

}
