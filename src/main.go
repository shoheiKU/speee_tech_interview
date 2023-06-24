package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server Start Up........")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
