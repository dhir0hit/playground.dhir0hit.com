package Controller

import (
	"net/http"
	"playground.dhir0hit.com/Controller/Home"
	"playground.dhir0hit.com/Controller/Playground"
	"playground.dhir0hit.com/Controller/Projects"
)

// TODO: LOAD TEMPLATE FROM VIEWS

func AppEntry() {
	http.HandleFunc("/", Home.Landing)
	http.HandleFunc("/home/", Home.Constructor)
	http.HandleFunc("/projects/", Projects.Constructor)
	http.HandleFunc("/playground/", Playground.Constructor)
	fs := http.FileServer(http.Dir("./WebApp/Public"))
	http.Handle("/Public/", http.StripPrefix("/Public", fs))
}
