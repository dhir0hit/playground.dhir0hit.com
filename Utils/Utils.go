package Utils

import (
	"errors"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Page Not Found")
	}
	return m[2], nil // The title is the second subexpression.
}

func ValidatePath(Views []string, pathArray []string) bool {
	for _, view := range Views {
		if strings.ToLower(view) == pathArray[0] {
			// IF Path Found in list
			return true
		}
	}
	// Else false
	return false
}

func GetPath(Views []string, w http.ResponseWriter, r *http.Request) ([]string, string, string) {
	var pathArray []string = strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	var controller string = pathArray[0]
	pathArray = pathArray[1:] // removing first element from array which is this controller name

	// Variable to store error caused by validation
	var validationError string

	// Checking if path Array has any element
	// Checking if path is valid
	if len(pathArray) > 0 {
		if !ValidatePath(Views, pathArray) {
			// Getting requested link and encoding it
			// to send as query in path
			var requestedPath string = url.PathEscape("/" + controller + "/" + strings.Join(pathArray, "/"))
			http.Redirect(w, r, "/"+controller+"/?error="+requestedPath, http.StatusFound)
			return nil, "error", validationError
		}
	} else {
		// Getting ALL Queries from URL
		query := r.URL.Query()
		err := query.Get("error")                  // getting error query from queries
		validationError, _ = url.PathUnescape(err) // Decoding / unEscaping url encoding
		pathArray = append(pathArray, "index")
	}
	return pathArray, controller, validationError
}
