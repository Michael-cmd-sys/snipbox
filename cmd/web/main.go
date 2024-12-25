package main

import (
  "net/http"
  "log"
  "flag"
  "os"
)

func main(){
  addr := flag.String("addr", ":5000", "HTTP Network Address");
  flag.Parse();

  infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime);
  errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Lshortfile);


  mux := http.NewServeMux();

  fileServer := http.FileServer(http.Dir("./ui/static"));

  mux.Handle("/static/", http.StripPrefix("/static", fileServer));
  mux.HandleFunc("/", home);
  mux.HandleFunc("/snippet/view", snippetView);
  mux.HandleFunc("/snippet/create", snippetCreate);

  srv := &http.Server{
    Addr: *addr,
    ErrorLog: errorLog,
    Handler: mux,
  }

  infoLog.Printf("Starting Server on %s", *addr);
  err := srv.ListenAndServe();
  errorLog.Fatal(err);
}
