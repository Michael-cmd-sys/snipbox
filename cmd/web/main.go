package main

import (
  "net/http"
  "log"
  "flag"
)

func main(){
  addr := flag.String("addr", ":5000", "HTTP Network Address");
  flag.Parse();

  mux := http.NewServeMux();

  fileServer := http.FileServer(http.Dir("./ui/static"));

  mux.Handle("/static/", http.StripPrefix("/static", fileServer));
  mux.HandleFunc("/", home);
  mux.HandleFunc("/snippet/view", snippetView);
  mux.HandleFunc("/snippet/create", snippetCreate);

  log.Printf("Starting server on %s", *addr);
  err := http.ListenAndServe(*addr, mux);
  log.Fatal(err);
}
