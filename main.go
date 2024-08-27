package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("starting the golang mongo with streams...")

	log.Fatal( http.ListenAndServe(":8080", nil))
	fmt.Println("starting the golang mongo server...")

	

}