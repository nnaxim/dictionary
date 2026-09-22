package main

import (
	"fmt"
	"net/http"
)

func (app *application) createWordHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new word")
}

func (app *application) showWordHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParams(r)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the details of word %d\n", id)
}
