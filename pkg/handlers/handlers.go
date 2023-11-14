package handlers

import (
	"go-sample-webserver/pkg/renders"
	"net/http"
)

func HomeHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	renders.RenderHtmlTemplate(res, "home.page.tmpl")
}

// This is the about page handler
func AboutHandler(res http.ResponseWriter, req *http.Request) {
	renders.RenderHtmlTemplate(res, "about.page.tmpl")
}
