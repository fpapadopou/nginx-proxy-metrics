package main

import (
	"flag"
	"greet/http"
	"log"
	nhttp "net/http"
	"strconv"
)

func main() {
	var host string
	var port int
	flag.StringVar(&host, "host", "127.0.0.1", "Hostname to which the app should listen to.")
	flag.IntVar(&port, "port", 8080, "Port to listen to.")

	flag.Parse()

	addr := host + ":" + strconv.Itoa(port)

	handler := http.New()

	log.Printf("HTTP server listening on: %v", addr)
	if err := nhttp.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}
