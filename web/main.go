package main

import (
	"net/http"

	"github.com/kasra1997/test/config"
	"github.com/kasra1997/test/pkg/handlers"
)

func main() {

	var app1 config.AppConfig
	app1.UseCache = false
	app1.Templatecache = handlers.CreateTemplateCache()
	handlers.Newtemplateconfset(&app1)
	repoo := handlers.Newrepo(&app1)
	handlers.Newhandler(repoo)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	http.ListenAndServe(":8002", nil)
}
