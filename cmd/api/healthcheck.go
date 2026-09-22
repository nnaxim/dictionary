package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: avilable")
	fmt.Fprintf(w, "envoronment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}
