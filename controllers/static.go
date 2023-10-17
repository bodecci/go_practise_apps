package controllers

import (
	"net/http"

	"github.com/bodecci/go_dev/view"
)

// create closure func that validates Templates at startup
func StaticHandler(tpl view.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}
