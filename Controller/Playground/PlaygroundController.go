package Playground

import (
	"fmt"
	"net/http"
	"playground.dhir0hit.com/Controller/Playground/CustomComponents"
	"playground.dhir0hit.com/Utils"
	"playground.dhir0hit.com/WebApp"
)

/*
PAGES

/playground/Home
*/

var Views []string = []string{"index", "components", "custom-components", "pages"}

func Constructor(w http.ResponseWriter, r *http.Request) {
	pathArray, title, validationError := Utils.GetPath(Views, w, r)

	fmt.Println(w, validationError)

	// if path array is not empty
	if pathArray != nil {
		// using switch case to connect path with right controller
		switch pathArray[0] {
		case "":
		case "home":
		case Views[0]:
			Index(w, r, title, pathArray)
			break

		case Views[1]:
			fmt.Println(w, "COMPONENTS")
			break

		case Views[2]:
			fmt.Println(w, "CUSTOM COMPONENTS")
			CustomComponents.Constructor(w, r, title, pathArray)
			break

		case Views[3]:
			fmt.Println(w, "PAGES")
			break

		default:
			fmt.Println(w, "NOT FOUND")
			break
		}
	}
}

func Index(w http.ResponseWriter, r *http.Request, title string, pathArray []string) {
	fmt.Println(w, "Playground > HOME")
	WebApp.Render(w, r, title, pathArray)
}
