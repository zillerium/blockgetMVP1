package main
 
import (
    "fmt"
    "html"
    "log"
    "net/http"
)
 
func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })
 
    fmt.Println("listening on port 9090")

    log.Fatal(http.ListenAndServe(":9090", nil))
 
}
