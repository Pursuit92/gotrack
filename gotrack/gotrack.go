package main

import (
	"net/http"
	"github.com/Pursuit92/gotrack"
	"log"
)

func main() {
	h := gotrack.NewHandler("15m","5m")
	http.Handle("/",h)
	log.Fatal(http.ListenAndServe(":8080",nil))
}
