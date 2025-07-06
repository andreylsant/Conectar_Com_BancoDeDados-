package main

import (
	"log"
	"net/http"

	"github.com/andrylsant/router"
)

func main() {
	router.HandelRouter()
	log.Fatal(http.ListenAndServe(":8000", nil))
}
