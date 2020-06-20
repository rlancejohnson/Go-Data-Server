package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/", TitlePage)
    http.ListenAndServe(GetPort(":8080"), nil)
}

func GetPort(default string) string {
    port := os.Getenv("PORT")
    if port != "" {
        return ":" + port
    } else {
        return default
    }
}

func TitlePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Go Data Server!")
}