package Projects

import (
	"fmt"
	"net/http"
	"playground.dhir0hit.com/Utils"
	"playground.dhir0hit.com/WebApp"
)

var Views []string = []string{"index"}

func Constructor(w http.ResponseWriter, r *http.Request) {
	pathArray, title, validationError := Utils.GetPath(Views, w, r)

	fmt.Println(validationError)

	// if path array is not empty
	if pathArray != nil {
		// using switch case to connect path with right controller
		switch pathArray[0] {
		case "":
		case "home":
		case Views[0]:
			Index(w, r, title, pathArray)
			break
		}
	}
}

func Index(w http.ResponseWriter, r *http.Request, title string, pathArray []string) {
	fmt.Println("PROJECTS > HOME ")
	WebApp.Render(w, r, title, pathArray)
}
