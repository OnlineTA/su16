package main

import (
    "fmt"
    "log"
    "net/http"
)

type Server struct {
  // Nothing to put here yet.
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    body := "Hello World\n"
    // Try to keep the same amount of headers
    w.Header().Set("Server", "gophr")
    w.Header().Set("Connection", "keep-alive")
    w.Header().Set("Content-Type", "text/plain")
    w.Header().Set("Content-Length", fmt.Sprint(len(body)))
    fmt.Fprint(w, body)
}

func main() {
  server := Server{}

  go func() {
    http.Handle("/", server)
      if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
      }
  }()
}
