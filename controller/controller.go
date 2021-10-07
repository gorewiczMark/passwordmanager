package controller

import (
	"fmt"
	"html/template"
	"net/http"

	p "passwordmanager/passentry"
	"passwordmanager/util/files"
	"passwordmanager/util/validations"

	"github.com/gorilla/mux"
)

var (
	passEntries    *p.Entries
	homeController home
	templates      map[string]*template.Template
)

// Startup start controller layer
func Startup(t map[string]*template.Template, entries *p.Entries) *mux.Router {
	passEntries = entries
	templates = t
	//homeController.homeTemplate = template["view.html"]
	//homeController.registerRoutes()
	// For images static handling
	//http.Handle("/img/", http.FileServer(http.Dir("public")))

	r := mux.NewRouter()
	// TODO Use after have a landing page otherwise fucks with the routing
	/*
		home := home{homeTemplate: templates["view.html"]}
		home.registerRoutes(r) */

	//r.HandleFunc("/", repoHandler) use for homepage
	repoRouter := r.PathPrefix("/repo/").Subrouter()
	repoRouter.HandleFunc("/", repoHandler)
	repoRouter.HandleFunc("/home", repoHandler)
	repoRouter.HandleFunc("/new", newHandler)
	repoRouter.HandleFunc("/namesearch", nameSearchHandler)
	repoRouter.HandleFunc("/groupsearch", groupSearchHandler)
	repoRouter.HandleFunc("/groupsearch/{group}", groupSearchHandler)
	repoRouter.HandleFunc("/search", searchHandler)
	repoRouter.HandleFunc("/groups", groupsHandler)
	repoRouter.HandleFunc("/create", createHandler)

	return r
}

func repoHandler(writer http.ResponseWriter, request *http.Request) {
	html := templates["view.html"]
	err := html.ExecuteTemplate(writer, "_layout.html", passEntries)
	validations.Check(err)
}

func newHandler(writer http.ResponseWriter, request *http.Request) {
	html := templates["new.html"]
	err := html.ExecuteTemplate(writer, "_layout.html", nil)
	validations.Check(err)
}

func nameSearchHandler(writer http.ResponseWriter, request *http.Request) {
	html := templates["search.html"]
	err := html.ExecuteTemplate(writer, "_layout.html", "name")
	validations.Check(err)
}

func groupSearchHandler(writer http.ResponseWriter, request *http.Request) {
	groups := passEntries.Groups()
	keys := make([]string, 0, len(groups))

	for k := range groups {
		keys = append(keys, k)
	}
	html := templates["searchgroup.html"]
	err := html.ExecuteTemplate(writer, "_layout.html", keys)
	validations.Check(err)
}

func searchHandler(writer http.ResponseWriter, request *http.Request) {
	var search, searchType string
	if request.Method == http.MethodPost {
		searchType = request.FormValue("searchType")
		search = request.FormValue("search")
	}

	var values p.Entries
	switch searchType {
	case "name":
		values = passEntries.SearchByName(search)
		break
	case "group":
		values = passEntries.SearchByGroup(search)
		break
	default:
	}

	html := templates["view.html"]
	err := html.ExecuteTemplate(writer, "_layout.html", values)
	validations.Check(err)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	var e p.Entry
	if r.Method == http.MethodPost {
		url := r.FormValue("url")
		e.Url = fmt.Sprintf("http://%s ", url)
		e.Username = r.FormValue("username")
		e.Password = r.FormValue("password")
		e.Name = r.FormValue("name")
		e.Grouping = r.FormValue("grouping")
		e.Extra = r.FormValue("extra")
		if r.FormValue("favorite") == "true" {
			e.Fav = true
		} else {
			e.Fav = false
		}
	}

	newEntry := fmt.Sprintf("%s,%s,%s,%s,%s,%s,%t", e.Url, e.Username, e.Password, e.Extra, e.Name, e.Grouping, e.Fav)

	files.WritePasswordsToFile("data/pass.csv", newEntry)
	passEntries.Insert(e, passEntries)
	http.Redirect(w, r, "/repo/", http.StatusFound)
}

func groupsHandler(writer http.ResponseWriter, request *http.Request) {
	groups := passEntries.Groups()
	keys := make([]string, 0, len(groups))
	for k := range groups {
		keys = append(keys, k)
	}

	html := templates["view.html"]
	err := html.ExecuteTemplate(writer, "_layout.html", keys)
	validations.Check(err)
}
