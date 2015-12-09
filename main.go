package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ant0ine/go-json-rest/rest"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/ping", func(w rest.ResponseWriter, r *rest.Request) {
			w.WriteJson("ok")
		}),
	)
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(2)
	}

	api.SetApp(router)

	url := "0.0.0.0:8080"
	fmt.Printf("Starting API on %s", url)
	if err := http.ListenAndServe(url, api.MakeHandler()); err != nil {
		fmt.Printf("Failed to start API %s", err.Error())
		os.Exit(2)
	}
}
