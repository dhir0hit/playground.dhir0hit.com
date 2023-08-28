package CustomComponents

import (
	"errors"
	"fmt"
	"net/http"
	"playground.dhir0hit.com/Data"
	"playground.dhir0hit.com/WebApp"
	"strings"
)

func Constructor(w http.ResponseWriter, r *http.Request, title string, pathArray []string) {
	fmt.Println(pathArray)
	fmt.Println(title)

	var err error = nil
	var data []Data.Elements = nil

	if len(pathArray) > 1 {
		if pathArray[1] == "index" || pathArray[1] == "home" {
			Index(w, r, title, pathArray)
			return
		}

		data, err = Data.LoadComponents(pathArray)
		if err == nil {
			// if Data Loaded
			SubController(w, r, title, pathArray, data)
			return
		}
	} else {
		Index(w, r, title, pathArray)
		return
	}
	return
}

func Index(w http.ResponseWriter, r *http.Request, title string, pathArray []string) {
	fmt.Fprintf(w, "COMPONENTS")
}

func SubController(w http.ResponseWriter, r *http.Request, title string, pathArray []string, data []Data.Elements) error {
	var resultData []Data.ComponentInfo = nil
	// If there is any data available
	if len(data) >= 1 {

		if len(pathArray) <= 2 {
			// TODO: create page for frontend or backend
			fmt.Fprintln(w, pathArray[1])
			return nil
		}

		// Looping through Data
		// finding correct element to navigate to
		for _, element := range data {
			// Getting Name of Component
			/*
				--custom-components.json--
				3 {
				4-- "name": "navigation" --
				5   "components": [...]
			*/
			if element.Name == pathArray[2] {
				// If there is any components in current element
				if len(element.Components) >= 1 && len(pathArray) > 3 {

					// getting and or from filter which is in link
					filterArray := strings.Split(pathArray[3], "&")
					for _, component := range element.Components {
						// Main sorting
						// Finding which are match for search or filter

						// if there is filter added in link
						// or not
						if len(filterArray) > 1 {
							// TODO: If want to add other type of filter add here
							if component.Framework == filterArray[0] && component.Language == filterArray[1] {
								resultData = append(resultData, component)
							}
						} else {
							if component.Framework == pathArray[3] || component.Language == pathArray[3] {
								resultData = append(resultData, component)
							}
						}
					}

					// rendering components
					WebApp.RenderComponentsContainer(w, title, pathArray, resultData)
					return nil
				}

				WebApp.RenderTagContainer(w, r, title, pathArray, element)
				return nil
			}
			// TODO: create error here that tag not found

		}
	}
	return errors.New("internal server error")
}
