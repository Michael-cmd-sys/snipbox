package main

import (
  "database/sql"
  "net/http"
  "html/template"
  "log"
  "flag"
  "os"
  "time"

  "github.com/Michael-cmd-sys/snipbox/internal/models"
  "github.com/alexedwards/scs/mysqlstore"
  "github.com/alexedwards/scs/v2"
  "github.com/go-playground/form/v4"
  _ "github.com/go-sql-driver/mysql"
)

type application struct {
  errorLog *log.Logger
  infoLog *log.Logger
  snippets *models.SnippetModel
  templateCache map[string]*template.Template
  formDecoder *form.Decoder
  sessionManager *scs.SessionManager
}

func main(){
  addr := flag.String("addr", ":5000", "HTTP Network Address");
  dsn := flag.String("dsn", "web:ssapword@/snipbox?parseTime=true", "MySQL data source name")

  flag.Parse();

  infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime);
  errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Lshortfile);


  db, err := openDB(*dsn)
  if err != nil {
    errorLog.Fatal(err)
  }
  defer db.Close()

  templateCache, err := newTemplateCache()
  if err != nil {
    errorLog.Fatal(err)
  }

  formDecoder := form.NewDecoder()

  sessionManager := scs.New()
  sessionManager.Store = mysqlstore.New(db)
  sessionManager.Lifetime = 12 * time.Hour

  sessionManager.Cookie.Secure = true

  app := &application{
    errorLog: errorLog,
    infoLog: infoLog,
    snippets: &models.SnippetModel{DB: db},
    templateCache: templateCache,
    formDecoder: formDecoder,
    sessionManager: sessionManager,
  }

  srv := &http.Server{
    Addr: *addr,
    ErrorLog: errorLog,
    Handler: app.routes(),
  }

  infoLog.Printf("Starting Server on %s", *addr);
  err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem");
  errorLog.Fatal(err);

}

func openDB(dsn string) (*sql.DB, error) {
  db, err := sql.Open("mysql", dsn)
  if err != nil {
    return nil, err
  }
  if err = db.Ping(); err != nil {
    return nil, err
  }

  return db, nil
}
