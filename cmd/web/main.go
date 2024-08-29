package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// Define a new command-line flag with the name 'addr'
	addr := flag.String("addr", ":8080", "HTTP network address")
	flag.Parse()



	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Staritng server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
