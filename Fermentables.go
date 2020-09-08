package main

import "strings"

type Fermentable struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	Country string `json:"country"`
	Category string `json:"category"`
	Type string `json:"type"`
	DegreesLovibond float64 `json:"degreesLovibond"`
	PPG float64 `json:"ppg"`
}

func allFermentables() []Fermentable {
	fermentables := retrieveFermentables()

	return fermentables
}

func getFermentableByName(name string) Fermentable {
	fermentables := retrieveFermentablesNameKey()
	if fermentable, found := fermentables[strings.ToLower(name)]; found {
		return fermentable
	}
	return Fermentable{Id: 0}
}

func getFermentableById(id int32) Fermentable {
	fermentables := allFermentables()
	for _, f := range fermentables {
		if f.Id == id {
			return f
		}
	}
	return Fermentable{Id: 0}
}