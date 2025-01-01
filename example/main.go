package main

import (
	"fmt"
	"net/http"

	"github.com/gocroot/example/route"
)

func main() {
	route.URL()
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
