package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
    "errors"

    "im3shn/snippetbox/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

    snippets, err := app.snippets.Latest()
    if err != nil {
        app.serverError(w, r, err)
        return
    }

    app.render(w, r, http.StatusOK, "home.tmpl.html", templateData{
        Snippets: snippets,
    })
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.PathValue("id"))
    if err != nil || id < 1 {
        http.NotFound(w, r)
        return
    }

    snippet, err := app.snippets.Get(id)
    if err != nil {
        if errors.Is(err, models.ErrNoRecord){
            http.NotFound(w, r)
            return
        }
        app.serverError(w, r, err)
        return
    }
    // fmt.Fprintf(w, "%+v", snippet)
    files := []string{
        "./ui/html/pages/base.tmpl.html",
        "./ui/html/partials/nav.tmpl.html",
        "./ui/html/pages/view.tmpl.html",
    }

    ts, err := template.ParseFiles(files...)
    if err != nil {
        app.serverError(w, r, err)
        return
    }

    data := templateData{
        Snippet: snippet,
    }

    err = ts.ExecuteTemplate(w, "base", data)
    if err != nil {
        app.serverError(w, r, err)
        return
    }

}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Display a form for creating a new snippet..."))
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
    title := "O Snail"
    content := "O Snail \nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
    expires := 7

    id, err := app.snippets.Insert(title, content, expires)

    if err != nil {
        app.serverError(w, r, err)
    }

    http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
