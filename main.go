package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"github.com/gorilla/mux"
)

var tpls *template.Template

//This function is called a beginning
func init() {
	var err error
	tpls, err = template.ParseFiles("assests/templates/index.html",
		"assests/templates/product.html","assests/templates/contact.html","assests/templates/service.html")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("File Parsed")

}

//Home page handler
func indexHandle(w http.ResponseWriter, r *http.Request) {

		err := tpls.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			log.Fatalln(err)
		}
}
func productHandler(w http.ResponseWriter,r *http.Request)  {
	err:=tpls.ExecuteTemplate(w,"product.html",nil)
	if err != nil {
		panic(err)
	}
}
func serviceHandler(w http.ResponseWriter,r *http.Request)  {
	err:=tpls.ExecuteTemplate(w,"service.html",nil)
	if err != nil {
		panic(err)
	}
}
func contactHandler(w http.ResponseWriter,r *http.Request)  {
	err:=tpls.ExecuteTemplate(w,"contact.html",nil)
	if err != nil {
		panic(err)
	}
}
func main() {
	mx:=mux.NewRouter()
	mx.HandleFunc("/", indexHandle)
	mx.HandleFunc("/product/",productHandler)
	mx.HandleFunc("/service/",serviceHandler)
	mx.HandleFunc("/contact/",contactHandler)
	mx.PathPrefix("/assests/").Handler(http.StripPrefix("/assests", http.FileServer(http.Dir("./assests"))))
	err := http.ListenAndServe(":80", mx)
	if err != nil {
		log.Fatalln(err)
	}
}