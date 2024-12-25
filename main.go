package main

import (
  "net/http"
  "log"
  "strconv"
  "fmt"
)

func home(w http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/" {
    http.NotFound(w, r);
    return;
  }

  w.Write([]byte("Hello from the home page..."));
}

func snippetView(w http.ResponseWriter, r *http.Request){
  id, err := strconv.Atoi(r.URL.Query().Get("id"));
  if err != nil || id < 1 {
    http.NotFound(w, r);
    return;
  }
  fmt.Fprintf(w, "Display a specific snippet with ID %d...", id);
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
  if r.Method != "POST" {
    w.Header().Set("Allow", http.MethodPost);
    http.Error(w, "Mehtod not allowed...", http.StatusMethodNotAllowed);
    return;
  }

  w.Write([]byte("Hello from the snippet create path"));
}

func main(){
  mux := http.NewServeMux();
  mux.HandleFunc("/", home);
  mux.HandleFunc("/snippet/view", snippetView);
  mux.HandleFunc("/snippet/create", snippetCreate);

  log.Println("Starting server on :5000");
  err := http.ListenAndServe(":5000", mux);
  log.Fatal(err);
}
