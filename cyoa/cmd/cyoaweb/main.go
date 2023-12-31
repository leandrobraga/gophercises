package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/leandrobraga/gophercises/cyoa"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the CYOA web application on")
	filename := flag.String("file", "gopher.json", "the JSON file with the CYOA history")
	flag.Parse()
	fmt.Printf("Using the story in %s\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonSotry(f)
	if err != nil {
		panic(err)
	}

	h := cyoa.NewHandler(story, nil)
	fmt.Printf("Starting the server port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
