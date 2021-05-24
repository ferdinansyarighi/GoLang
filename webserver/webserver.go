package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

// Web Server using GO
// run this and open localhost:8080/anything to test
// learned by Muhammad Ferdinansyah Arighi
// source : https://golang.org/doc/articles/wiki/