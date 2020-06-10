package main

import (
        "fmt"
        "log"
        "net/http"
        "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
        log.Print("return-token: received a request")
        token := os.Getenv("TOKEN")
        if token == "" {
                token = "Specify a token to be returned via the TOKEN env var"
        }
        fmt.Fprintf(w, token)
}

func main() {
        log.Print("return-token: starting server...")

        http.HandleFunc("/", handler)

        port := os.Getenv("PORT")
        if port == "" {
                port = "8080"
        }

        log.Printf("return-token: listening on port %s", port)
        log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}