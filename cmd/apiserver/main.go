package main

import (
  "log"
  "github.com/Vlad0siy/http-rest-api/internal/app/apiserver"
)

func main() {
  s := apiserver.new()
  if err := s.Start(), err != nil {
    log.Fatal(err)
  }
}
