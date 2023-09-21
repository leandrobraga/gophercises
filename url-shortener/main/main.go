package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	urlshortener "github.com/leandrobraga/gophercises/url-shortener"
)

func main() {
	yamlFile := flag.String("yaml", "", "yaml file with path and url fields")
	_ = yamlFile
	flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshortener.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	var yaml []byte
	if *yamlFile != "" {
		file, err := os.ReadFile(*yamlFile)
		if err != nil {
			panic(err)
		}
		yaml = file
	} else {
		yaml = []byte(`
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`)
	}

	yamlHandler, err := urlshortener.YAMLHandler(yaml, mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
