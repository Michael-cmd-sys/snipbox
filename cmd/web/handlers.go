package main 

import (
  "fmt"
  "net/http"
  "html/template"
  "log"
  "strconv"
)

func home(w http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/" { 
    http.NotFound(w, r);
    return;
  }

  files := []string{
    "./ui/html/base.tmpl.html",
    "./ui/html/pages/home.tmpl.html",
  }

  ts, err := template.ParseFiles(files...);
  if err != nil {
    log.Println(err.Error());
    http.Error(w, "Internal server error", http.StatusInternalServerError);
  }

  err = ts.ExecuteTemplate(w, "base", nil);
  if err != nil {
    log.Println(err.Error());
    http.Error(w, "Internal Server Error", http.StatusInternalServerError);
  }
}

func snippetView(w http.ResponseWriter, r *http.Request) {
  id, err := strconv.Atoi(r.URL.Query().Get("id"));
  if err != nil || id < 1 {
    http.NotFound(w, r);
    return;
  }

  fmt.Fprintf(w, "Displaying snippet for ID %d", id);
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
    w.Header().Set("Allow", http.MethodPost);
    http.Error(w, "Method not allowed...", http.StatusMethodNotAllowed);
    return;
  }

  w.Write([]byte("Create new snippet..."));
}
