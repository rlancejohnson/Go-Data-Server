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

func GetPort(defaultPort string) string {
    port := os.Getenv("PORT")
    if port != "" {
        return ":" + port
    } else {
        return defaultPort
    }
}

func TitlePage(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path[1:] != "" {
        file, err := ioutil.ReadFile(r.URL.Path[1:])
        printErr(err)

        fmt.Fprintf(w, string(file))

    } else {
        fmt.Fprintf(w, "Please append the url with a file name.")
    }
}

func printErr(err error) {
    if err != nil {
        fmt.Println(err)
    }
}