package main

import (
  "net/http"
  "log"
)

func main(){
  mux := http.NewServeMux();
  mux.HandleFunc("/", home);
  mux.HandleFunc("/snippet/view", snippetView);
  mux.HandleFunc("/snippet/create", snippetCreate);

  log.Println("Starting server on :5000");
  err := http.ListenAndServe(":5000", mux);
  log.Fatal(err);
}
