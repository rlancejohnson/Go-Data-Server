package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = ":8080"
    }

    http.HandleFunc("/", TitlePage)
    http.ListenAndServe(port, nil)
}

func TitlePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Go Data Server!")
}