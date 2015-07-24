package main

import (
  "os"
	"net/http"
  "fmt"

	"github.com/ant0ine/go-json-rest/rest"
)

func main() {
  handler := rest.ResourceHandler{}

	err := handler.SetRoutes(
		&rest.Route{"GET", "/ping", func(w rest.ResponseWriter, r *rest.Request) {
      w.WriteJson("ok")
    }},
  )
	if err != nil {
    fmt.Printf(err.Error())
		os.Exit(2)
	}

  url := "0.0.0.0:8080"
	fmt.Printf("Starting API on %s", url)
	if err := http.ListenAndServe(url, &handler); err != nil {
		fmt.Printf("Failed to start API %s", err.Error())
		os.Exit(2)
	}
}
