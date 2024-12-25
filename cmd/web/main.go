package main

import (
  "net/http"
  "log"
  "flag"
  "os"
)

type application struct {
  errorLog *log.Logger
  infoLog *log.Logger
}

func main(){
  addr := flag.String("addr", ":5000", "HTTP Network Address");
  flag.Parse();

  infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime);
  errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Lshortfile);

  app := &application{
    errorLog: errorLog,
    infoLog: infoLog,
  }

  srv := &http.Server{
    Addr: *addr,
    ErrorLog: errorLog,
    Handler: app.routes(),
  }

  infoLog.Printf("Starting Server on %s", *addr);
  err := srv.ListenAndServe();
  errorLog.Fatal(err);
}
