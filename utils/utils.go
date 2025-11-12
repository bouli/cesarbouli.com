package utils

import (
	"log"
	"net/http"
	"os"
)

func GetHttpProtocolString(r *http.Request) string {
	if r.TLS != nil {
		return "https"
	} else {
		return "http"
	}
}

func GetServerPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT is not setted. Assuming 80 as default.")
		return "80"
	} else {
		return port
	}
}
