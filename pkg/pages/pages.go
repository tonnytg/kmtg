package pages

import (
	"github.com/tonnytg/kmtg/domain/pods"
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	// call template home and parse data from PODs
	p := pods.Pods{}
	// send struct to html called home
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	tmpl.Execute(w, p)
}

func Listen() {
	// delivery home with template parse blob template with home.html
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
