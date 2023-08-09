package main

import (
	"log"
	"net/http"

	"hello/components"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func main() {
	app.Route("/", &components.AppControl{})
	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Name:        "Tab Title",
		Description: "Give browser tab a meaningful title",
	})

	log.Println("Listening on http://:8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
