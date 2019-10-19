package main

import (
	"fmt"
	"net/http"

	urlShortener "github.com/ttimt/urlshort-gophercises"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlShortener-godoc": "https://godoc.org/github.com/gophercises/urlShortener",
		"/yaml-godoc":         "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlShortener.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml := `
- path: /urlShortener
  url: https://github.com/gophercises/urlShort
- path: /urlShortener-final
  url: https://github.com/gophercises/urlShort/tree/solution
`
	yamlHandler, err := urlShortener.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	_ = http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintln(w, "Hello, world!")
}
