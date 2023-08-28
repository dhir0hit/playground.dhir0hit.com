package WebApp

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"playground.dhir0hit.com/Data"
	"strings"
)

type Page struct {
	Title    string
	QuickNav template.HTML
	Body     template.HTML
}

//func (p *Page) save() error {
//	filename := p.Title + ".txt"
//	return os.WriteFile(filename, byte(p.Body), 0600)
//}

func _loadPage(title string, pathArray []string) (*Page, error) {
	var viewPath string = "./WebApp/Views/" + title + "/" + strings.Join(pathArray, "/") + ".html"
	body, err := os.ReadFile(viewPath)
	if err != nil {
		//fmt.Println(err)
		return nil, err
	}
	quickNav := addPath(title, pathArray)

	return &Page{Title: title, QuickNav: template.HTML(quickNav), Body: template.HTML(body)}, nil
}

func _renderLayout(w http.ResponseWriter, p *Page) {
	//err := templates.ExecuteTemplate(w, tmpl+".html", p)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//}

	var LayoutPath string = "./WebApp/Views/layout.html"

	w.Header().Set("Content-Type", "text/html")
	t, err := template.ParseFiles(LayoutPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func addPath(title string, pathArray []string) string {
	var path string
	path += "<div class='quick-nav'>"
	path += "<a href='/" + title + "'>" + title + "</a>"

	var subPath string = "/" + title
	for _, _path := range pathArray {
		if _path == "index" {
			break
		}
		subPath += "/" + _path
		path += " > "
		path += "<a href='" + subPath + "'>" + _path + "</a>"
	}
	path += "></div>"

	return path
}

func Render(w http.ResponseWriter, r *http.Request, title string, pathArray []string) {
	fmt.Println("Rendering Page...")
	w.Header().Set("Content-Type", "text/html")
	page, err := _loadPage(title, pathArray)
	fmt.Println(pathArray)
	if err != nil {
		// redirect to error page
		return
	}
	_renderLayout(w, page)
}

func RenderTagContainer(w http.ResponseWriter, r *http.Request, title string, pathArray []string, data Data.Elements) {

	var Body string
	Body += "<div class='lang small container'>"
	for _, element := range data.Components {
		var Language string = element.Language

		if element.Framework != "" {
			if strings.ToLower(Language) == "javascript" {
				Language = " JS"
			}
			if strings.ToLower(Language) == "typescript" {
				Language = " TS"
			}
		}
		Body += "<a class='sub-container'><i class='fa-brands fa-2xl fa-react'></i><h3>" + element.Framework + Language + "</h3></a>"
	}
	Body += "</div>"

	quickNav := addPath(title, pathArray)
	page := &Page{Title: title, QuickNav: template.HTML(quickNav), Body: template.HTML(Body)}
	_renderLayout(w, page)
}

func RenderComponentsContainer(w http.ResponseWriter, title string, pathArray []string, data []Data.ComponentInfo) {
	// Todo: limit 3 per page
	var Body string

	// Looping on data
	for _, element := range data {
		var lang string = element.Language
		if strings.ToLower(element.Language) == "javascript" {
			lang = "JS"
		} else if strings.ToLower(element.Language) == "typescript" {
			lang = "TS"
		}
		// TODO: Load from component.html
		Body += "<div class='online-pen-info-container'>" +
			"<h1>" + element.Framework + "-" + lang + "</h1>" +
			"<h3>Created by: " + element.Owner + "</h3>" +
			"<p>More info</p>" +
			"<iframe src='" + element.Link + "'" +
			"style='width: 100%; height: 500px; border: 0;background:#000;'" +
			"></iframe>" +
			"</div>"
	}

	quickNav := addPath(title, pathArray)
	page := &Page{Title: title, QuickNav: template.HTML(quickNav), Body: template.HTML(Body)}

	_renderLayout(w, page)
}

func RenderComponentDetail(w http.ResponseWriter, title string, pathArray []string, data []Data.ComponentInfo) {
	// TODO: maybe ill use it
	// limit 3 per page
	fmt.Println(data)
}
