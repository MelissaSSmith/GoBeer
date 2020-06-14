package main

import (
	"fmt"
	"strings"
)

type Fermentable struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	Country string `json:"country"`
	Category string `json:"category"`
	Type string `json:"type"`
	DegreesLovibond float32 `json:"degreesLovibond"`
	PPG float32 `json:"ppg"`
}

func allFermentables() []Fermentable {
	fermentables := retrieveFermentables()

	return fermentables
}

func getFermentableByName(name string) Fermentable {
	fermentables := allFermentables()

	for _, f := range fermentables {
		if strings.ToLower(f.Name) == strings.ToLower(name) {
			return f
		}
	}
	return Fermentable{Id: 0}
}

func getFermentableById(id int32) Fermentable {
	fermentables := allFermentables()
	fmt.Println("Here")
	for _, f := range fermentables {
		if f.Id == id {
			return f
		}
	}
	return Fermentable{Id: 0}
}