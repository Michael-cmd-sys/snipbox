package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
  "strings"
  "unicode/utf8"

	"github.com/Michael-cmd-sys/snipbox/internal/models"
  "github.com/Michael-cmd-sys/snipbox/internal/validator"

	"github.com/julienschmidt/httprouter"
)


type snippetCreateForm struct {
  FieldErrors map[string]string
  validator.Validator
  Title string
  Content string
  Expires int
}
func (app *application) home(w http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/" { 
    app.notFound(w);
    return;
  }

  snippets, err := app.snippets.Latest()
  if err != nil {
    app.serverError(w, err)
    return
  }

  data := app.newTemplateData(r)
  data.Snippets = snippets

   app.render(w, http.StatusOK, "home.tmpl.html", data)
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
  params := httprouter.ParamsFromContext(r.Context())

  id, err := strconv.Atoi(params.ByName("id"))
  if err != nil || id < 1 {
    app.notFound(w);
    return;
  }

  snippet, err := app.snippets.Get(id)
  if err != nil {
    if errors.Is(err, models.ErrNoRecord) {
      app.notFound(w)
    } else {
      app.serverError(w, err)
    }

    return
  }
  data := app.newTemplateData(r)
  data.Snippet = snippet

  app.render(w, http.StatusOK, "view.tmpl.html", data)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
  data := app.newTemplateData(r)

  data.Form = snippetCreateForm{
    Expires: 365,
  }
  app.render(w, http.StatusOK, "create.tmpl.html", data)
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
  err := r.ParseForm()
  if err != nil {
    app.clientError(w, http.StatusBadRequest)
    return 
  }

  expires, err := strconv.Atoi(r.PostForm.Get("expires"))
  if err != nil {
    app.clientError(w, http.StatusBadRequest)
    return
  }

  form := snippetCreateForm{
    Title: r.PostForm.Get("title"),
    Content: r.PostForm.Get("content"),
    Expires: expires,
  }

  form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be blank.")
  form.CheckField(validator.MaxChars(form.Title, 100), "title", "This field cannot be more than 100 characters long.")
  form.CheckField(validator.NotBlank(form.Content), "content", "This field cannot be blank.")
  form.CheckField(validator.PermittedInt(form.Expires, 1, 7, 365), "expires", "This field must equal 1, 7 or 365")

  if !form.Valid() {
    data := app.newTemplateData(r)
    data.Form = form
    app.render(w, http.StatusUnprocessableEntity, "create.tmpl.html", data)
    return
  }

  id , err := app.snippets.Insert(form.Title, form.Content, expires)
  if err != nil {
    app.serverError(w, err)
    return
  }

  http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
