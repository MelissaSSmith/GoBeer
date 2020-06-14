package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

var fermentableFile = "./data/fermentables.json"
var srmHexFile = "./data/srm_hex.json"

func retrieveFermentables() []Fermentable {
	data, err := ioutil.ReadFile(fermentableFile)
	if err != nil {
		fmt.Print(err)
	}
	var obj []Fermentable

	err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Print("error:", err)
	}

	return obj
}

func retrieveFermentablesNameKey() map[string]Fermentable {
	fermentables := retrieveFermentables()
	fermentableMap := make(map[string]Fermentable)

	for _, f := range fermentables {
		fermentableMap[strings.ToLower(f.Name)] = f
	}
	return fermentableMap
}

func retrieveSrmHexValues() map[int]SrmHex {
	data, err := ioutil.ReadFile(srmHexFile)
	if err != nil {
		fmt.Print(err)
	}
	var obj []SrmHex
	srmHexMap := make(map[int]SrmHex)

	err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Print("error:", err)
	}
	for _, s := range obj {
		srmHexMap[s.SrmKey] = s
	}

	return srmHexMap
}