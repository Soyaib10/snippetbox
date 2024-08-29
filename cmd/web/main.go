package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// application struct application-wide dependencies
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
    // Define a new command-line flag
    addr := flag.String("addr", ":8080", "HTTP network address")
    flag.Parse()

    // Creating custom loggers
    infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
    errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

    // Initialize a new instance of application containing the dependencies.
    app := &application{
        errorLog: errorLog,
        infoLog: infoLog,
    }

    mux := http.NewServeMux()
    mux.HandleFunc("/", app.home)
    mux.HandleFunc("/snippet", app.showSnippet)
    mux.HandleFunc("/snippet/create", app.createSnippet)

    fileServer := http.FileServer(http.Dir("./ui/static/"))
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

    srv := &http.Server{
        Addr:     *addr,
        ErrorLog: errorLog,
        Handler:  mux,
    }

    infoLog.Printf("Starting server on %s", *addr)
    err := srv.ListenAndServe()
    if err != nil {
        errorLog.Fatalf("Server failed: %v", err)
    }
}

