package main

import (
	"fer/src/infraestructure/routes"
	"net/http"
)

func main() {
	r := routes.InitRouter()

	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
}
