package main

import (
	"encoding/csv"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"passwordmanager/controller"
	p "passwordmanager/passentry"
	"passwordmanager/util/files"
	"passwordmanager/util/validations"
	"strconv"
	"strings"
)

// TODO: Variables for now look at removing
var passEntries p.Entries
var templates map[string]*template.Template

func getPasswords(filename string) [][]string {
	// os.Open(filename)
	flags := os.O_RDONLY
	file, err := os.OpenFile(filename, flags, os.FileMode(0666))
	if os.IsNotExist(err) {
		return nil
	}
	validations.Check(err)
	defer file.Close()

	r := csv.NewReader(file)
	r.Read()                  //read head of csv file and toss
	lines, err := r.ReadAll() //read rest of file
	validations.Check(err)
	return lines
}

func parseLines(lines *[][]string) p.Entries {
	ret := make([]p.Entry, len(*lines))
	for i, line := range *lines {
		favorite, _ := strconv.ParseBool(line[6])

		ret[i] = p.Entry{
			Url:      line[0],
			Username: strings.TrimSpace(line[1]),
			Password: strings.TrimSpace(line[2]),
			Extra:    strings.TrimSpace(line[3]),
			Name:     strings.TrimSpace(line[4]),
			Grouping: strings.TrimSpace(line[5]),
			Fav:      favorite,
		}
	}
	return ret
}

const (
	baseTemp    = "templates/"
	baseContent = baseTemp + "/content/"
)

// prepApp - helper function to read the data add create a [] of entries
func prepApp() {
	lines := files.GetPasswordsFromFile("data/pass.csv")
	passEntries = parseLines(&lines)
}

// main
func main() {
	/* Messing arround with flags
	   boolPtr := flag.Bool("exit", false, "exit")
	   	flag.Parse()

	   	if len(os.Args) > 1 {
	   		if *boolPtr {
	   			fmt.Println("Exiting because of flag")
	   			os.Exit(0)
	   		}
	   	} */
	prepApp()
	templates = populateTemplates()

	r := controller.Startup(templates, &passEntries)
	err := http.ListenAndServeTLS("localhost:8080", "cert.pem", "key.pem", r)
	if err != nil {
		log.Fatal(err)
	}
}

func populateTemplates() map[string]*template.Template {
	result := make(map[string]*template.Template)
	const basePath = "templates"
	layout := template.Must(template.ParseFiles(basePath + "/_layout.html"))
	template.Must(layout.ParseFiles(basePath+"/_header.html", basePath+"/_footer.html"))
	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}
	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}
	for _, fi := range fis {
		f, err := os.Open(basePath + "/content/" + fi.Name())
		if err != nil {
			panic("Failed to open template '" + fi.Name() + "'")
		}
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to read content from file '" + fi.Name() + "'")
		}
		f.Close()
		tmpl := template.Must(layout.Clone())
		_, err = tmpl.Parse(string(content))
		if err != nil {
			panic("Failed to parse contents of '" + fi.Name() + "' as template")
		}
		result[fi.Name()] = tmpl
	}
	return result
}
