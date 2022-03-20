package main
import (
 "fmt"
 "net/http"
)

// handler - Prints out "OK Let's Go" and URL path on the screen.
func handler(writer http.ResponseWriter, request *http.Request) {
 fmt.Fprintf(writer, "OK Let's Go, %s!", request.URL.Path[1:])
}

func main() {
 http.HandleFunc("/", handler)
 http.ListenAndServe(":8080", nil)
}