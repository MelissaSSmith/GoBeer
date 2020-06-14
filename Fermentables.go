package main

type Fermentable struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	Country string `json:"country"`
	Category string `json:"category"`
	Type string `json:"type"`
	DegreesLovibond float32 `json:"degreesLovibond"`
	PPG float32 `json:"ppg"`
}
