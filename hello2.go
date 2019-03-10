package main

import(
  "os"
  "fmt"
  "net/http"
)

func main() {
  hostname := os.Getenv("HOSTNAME")
  if hostname == "" {
    hostname = "World!"
  }

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "New Hello, %s", hostname)
  })

  http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "2.0.0")
  })

  http.ListenAndServe(":3000", nil)

}
