package main

import (
  "fmt"
  "net/http"
  "os"

	"github.com/sirupsen/logrus"
)

func health(w http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(w, "ok\n")
}

func main() {
  fs := http.FileServer(http.Dir("./static"))
  http.Handle("/", fs)
  http.HandleFunc("/health", health)

  port := "8090"
  if len(os.Getenv("PORT")) > 0 {
    port = os.Getenv("PORT")
  }

  logrus.Infof("Listening on :%s...", port)
  err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
  if err != nil {
    logrus.WithError(err).Fatal("Failed serving")
  }
}
