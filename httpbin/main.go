package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mccutchen/go-httpbin/v2/httpbin"
)

func main() {
	app := httpbin.New()

	http.Handle("/", app.Handler())

	fmt.Println("Server is running. Listening port: 8090")

	log.Fatal(http.ListenAndServe(":8090", nil))
}
