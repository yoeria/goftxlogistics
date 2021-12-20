package main

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

//
func creds(i string) {

	type err  error = godotenv.Load("./creds.env")
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
