package Data

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
)

type ComponentInfo struct {
	Language  string   `json:"language"`
	Framework string   `json:"framework"`
	Owner     string   `json:"owner"`
	Mail      string   `json:"mail"`
	Github    string   `json:"github"`
	Info      string   `json:"info"`
	Link      string   `json:"link"`
	HowTo     []string `json:"how-to"`
}

type Elements struct {
	Name        string          `json:"name"`
	BannerImage string          `json:"banner-image"`
	Tags        string          `json:"tags"`
	Components  []ComponentInfo `json:"components"`
}

type Data struct {
	Frontend []Elements `json:"frontend"`
	Backend  []Elements `json:"backend"`
}

func LoadComponents(pathArray []string) ([]Elements, error) {
	fileContent, err := os.ReadFile("./Data/" + strings.ToLower(pathArray[0]) + ".json")
	if err != nil {
		return nil, err
	}
	var data Data
	err = json.Unmarshal(fileContent, &data)
	if err != nil {
		return nil, err
	}

	if strings.ToLower(pathArray[1]) == "frontend" {
		return data.Frontend, nil
	}
	if strings.ToLower(pathArray[1]) == "backend" {
		return data.Backend, nil
	}
	return nil, errors.New("internal server error")

}
