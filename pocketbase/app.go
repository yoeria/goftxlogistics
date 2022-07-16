package pocketbase

import (
	"log"

	"github.com/pocketbase/pocketbase"
)

func Exec() {
	app := pocketbase.New()

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
