package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	log.Println("time-app: starting time app")
	timefmt := os.Getenv("TIME_FORMAT")
	switch timefmt {
	case "ANSIC":
		timefmt = time.ANSIC
	case "RFC822":
		timefmt = time.RFC822
	case "RFC822Z":
		timefmt = time.RFC822Z
	case "RFC1123":
		timefmt = time.RFC1123
	case "RFC1123Z":
		timefmt = time.RFC1123Z
	case "UNIX":
		timefmt = time.UnixDate
	default:
		if timefmt == "" {
			timefmt = time.RFC822Z
		}
	}
	log.Printf("time-app: format: %s", timefmt)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Print("req received")
		fmt.Fprintln(w, time.Now().Format(timefmt))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("time-app: listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
