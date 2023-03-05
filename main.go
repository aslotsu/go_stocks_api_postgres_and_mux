package main

import (
	"fmt"
	"go-pg13/router"
	"net/http"
)

func main() {

	r := router.Router()
	fmt.Println("Starting the server on port 8080...")

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println(err.Error())
	}

}
