package main

import "strings"

type Hop struct {
	Name string `json:"name"`
	HopId string `json:"hop_id"`
	Characteristics string `json:"characteristics"`
	Purpose string `json:"purpose"`
	Country string `json:"country"`
	AlphaAcidComposition string `json:"alpha_acid_composition"`
	BetaAcidComposition string `json:"beta_acid_composition"`
	CoHumuloneComposition string `json:"co_humulone_composition"`
	SeasonalMaturity string `json:"seasonal_maturity"`
	YieldAmount string `json:"yield_amount"`
	GrowthRate string `json:"growth_rate"`
	SusceptibleTo string `json:"susceptible_to"`
	Storability string `json:"storability"`
	TotalOilComposition string `json:"total_oil_composition"`
	StyleGuide []string `json:"style_guide"`
	Substitutes [] string `json:"substitutes"`
	ResistantTo string `json:"resistant_to"`
	MyrceneOilComposition string `json:"myrcene_oil_composition"`
	CaryophylleneOil string `json:"caryophyllene_oil"`
	FarneseneOil string `json:"farnesene_oil"`
	AlsoKnownAs string `json:"also_known_as"`
}

func allHops() []Hop {
	hops := retrieveHops()

	return hops
}

func getHopByName(name string) Hop {
	println(name)
	hops := retrieveHopsIdKey()
	if hop, found := hops[strings.ToLower(name)]; found {
		return hop
	}
	return Hop{HopId: ""}
}