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
	jsonFile := flag.String("json", "", "json file with path and url fields")

	flag.Parse()

	if *yamlFile != "" && *jsonFile != "" {
		fmt.Println("Only one option is accepted: yaml or json")
		os.Exit(1)
	}

	mux := defaultMux()

	var handler http.HandlerFunc

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshortener.MapHandler(pathsToUrls, mux)

	if *jsonFile != "" {
		var err error
		file, err := os.ReadFile(*jsonFile)
		if err != nil {
			panic(err)
		}
		handler, err = urlshortener.JSONHandler(file, mapHandler)
		if err != nil {
			panic(err)
		}

	}

	if *jsonFile == "" {
		var yaml []byte
		var err error
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
		handler, err = urlshortener.YAMLHandler(yaml, mapHandler)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", handler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
