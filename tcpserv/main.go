package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

func dog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "dog dog doggy")
}

func cat(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "cat cat cat")
}

var tmpl *template.Template

func main() {
	tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))

	mux := httprouter.New()
	mux.GET("/", index)
	mux.GET("/about", about)
	mux.GET("/contact", contact)
	mux.GET("/apply", applyProcess)
	mux.POST("/apply", applyProcess)
	mux.GET("/user/:name", user)
	mux.GET("/blog/:category/:article", blogRead)
	mux.POST("/blog/:category/:article", blogWrite)
	http.ListenAndServe(":8080", mux)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tmpl.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err)
}

func about(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tmpl.ExecuteTemplate(w, "about.gohtml", nil)
	HandleError(w, err)
}

func contact(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tmpl.ExecuteTemplate(w, "apply.gohtml", nil)
	HandleError(w, err)
}

func applyProcess(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tmpl.ExecuteTemplate(w, "applyProcess.gohtml", nil)
	HandleError(w, err)
}

func user(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "USER %s!\n", ps.ByName("name"))
}

func blogRead(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "READ CATEGORY, %s!\n", ps.ByName("category"))
	fmt.Fprintf(w, "READ ARTICLE, %s!\n", ps.ByName("article"))
}

func blogWrite(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "WRITE CATEGORY, %s!\n", ps.ByName("category"))
	fmt.Fprintf(w, "WRITE ARTICLE, %s!\n", ps.ByName("article"))
}

func update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "UPDATE, %s!\n", ps.ByName("name"))
}

func trailing(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "TRAILING, %s!\n", ps.ByName("name"))
}

func HandleError(w http.ResponseWriter, e error) {
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		log.Println(e)
	}
}
