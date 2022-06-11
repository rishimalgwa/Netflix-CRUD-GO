package main

import (
	"fmt"
	"hello/23mongodbapi/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000...")
}
