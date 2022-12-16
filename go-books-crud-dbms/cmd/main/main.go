package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Print("Server is running on port 4000")
	if err := http.ListenAndServe(":4000", nil); err != nil {
		log.Fatal(err)
	}
}
