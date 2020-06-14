package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func retrieveSrmHexValues() []SrmHex {
	data, err := ioutil.ReadFile(srmHexFile)
	if err != nil {
		fmt.Print(err)
	}
	var obj []SrmHex

	err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Print("error:", err)
	}

	return obj
}