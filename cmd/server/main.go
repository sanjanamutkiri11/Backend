package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting Cloud Autoscaler Backend...")
	// TODO: Initialize config, DB, and router
	log.Fatal(http.ListenAndServe(":8080", nil))
}
