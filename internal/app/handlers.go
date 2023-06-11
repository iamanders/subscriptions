package app

import (
	"net/http"
)

// GET /
func (app *App) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("HELLO?"))
	}
}
