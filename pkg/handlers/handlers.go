package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/kasra1997/test/config"
	"github.com/kasra1997/test/models"
)

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	strmap := make(map[string]string)
	strmap["tt"] = "this came from template value trasmision "
	RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: strmap,
	})
}

var app *config.AppConfig

//Newtemplateconfset set the initial template renders for app and it make us to dont render it many times
func Newtemplateconfset(a *config.AppConfig) {
	app = a
}

////////////////////////////////////////////////////////////////////////////////////////////////
//repo is the repository for the handlers
var Repo *Repository

type Repository struct {
	app *config.AppConfig
}

// create new repository
func Newrepo(a *config.AppConfig) *Repository {
	return &Repository{
		app: a,
	}
}

//set the repository for the handlers
func Newhandler(r *Repository) {
	Repo = r
}

//////////RRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRRender render render render render render render render render render part is here
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var t map[string]*template.Template
	if app.UseCache {
		t = app.Templatecache
	} else {
		t = CreateTemplateCache()
	}

	if len(t) == 0 {
		fmt.Println("lenth of t (createtemplatecache,app.templatecache,config.app.templatecache) is zero")
	}

	buff := new(bytes.Buffer)
	tm, ok := t[tmpl]
	if !ok {
		buff.WriteString(fmt.Sprintf("%s\n", tmpl))
		return
	}

	tm.Execute(buff, td)
	_, err := buff.WriteTo(w)
	if err != nil {
		panic(err)
	}

	/*parsedTemplate, _ := template.ParseFiles("./layout/_default/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("\n fuck you for this error , code :fart \n", err)
		return*/
}

var functions = template.FuncMap{
	//"foobar": func(a string) string {
	//	return a + "-foobar"
	//},
}

func CreateTemplateCache() map[string]*template.Template {
	mycache := map[string]*template.Template{}

	pagess, err := filepath.Glob("./layout/_default/*.page.html")
	if err != nil {
		fmt.Println("render template t/base errrrrrrrrrrrrrrrrrrrrrrrrrrrrrrr")
		return mycache
	}

	for _, page := range pagess {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			fmt.Println("render template t/base errrrrrrrrrrrrrrrrrrrrrrrrrrrrrrr222")
			return mycache
		}
		ts, err = ts.ParseGlob("layout/_default/baseof.html")
		if err != nil {
			fmt.Println(err)
			return mycache
		}
		mycache[name] = ts

	}
	return mycache
}
