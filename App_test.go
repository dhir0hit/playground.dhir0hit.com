package main

import (
	"fmt"
	"net/http"
	"testing"
)

type Views map[string][]string

func TestAppEntry(t *testing.T) {
	fmt.Println("")
	fmt.Println("Specified URL Test")
	fmt.Println("-----------------------------------------------------------------------")
	views := Views{
		"":           []string{""},
		"Home":       []string{"", "index"},
		"Playground": []string{"", "index", "components, pages"},
		"Project":    []string{"", "index"},
	}
	url := "http://localhost:8080/"

	for controller, views := range views {

		url += controller + "/"

		for _, view := range views {
			url += view + "/"

			resp, err := http.Get(url)
			if err != nil {
				t.Fatalf("Failed to GET URL %s: %v", url, err)
			}
			if resp.StatusCode != http.StatusOK {
				t.Errorf("Expected Status Code %d, but got %d", http.StatusOK, resp.StatusCode)
			}
		}
	}

}
