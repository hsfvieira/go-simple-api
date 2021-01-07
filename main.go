package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/users/", handleUsers)
	http.ListenAndServe(":8080", nil)
}