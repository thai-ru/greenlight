package main

import (
	"fmt"
	"net/http"
)

// healtcheckHandler writes a human readable text with information about the application status, environment and version
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: OK")
	fmt.Fprintf(w, "version: %s\n", version)
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
}
