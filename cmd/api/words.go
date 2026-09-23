package main

import (
	"dictionary-api/internal/data"
	"fmt"
	"net/http"
	"time"
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

	data := data.Words{
		ID:        1,
		CreatedAt: time.Now(),
		Title:     "Hello",
		Example:   "",
		Synonims:  []string{"Hi", "Hey", "Hi there", "Hey there"},
		Version:   1,
	}

	if id != data.ID {
		app.logger.Println("ID unexists")
		http.Error(w, "ID unexists", http.StatusNotFound)
		return
	}

	err = app.writeJson(w, http.StatusOK, data, nil)

	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
