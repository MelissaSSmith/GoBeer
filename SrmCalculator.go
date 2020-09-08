package main

import (
	"math"
	"strings"
)

type SrmRequest struct {
	GrainBill []RecipeGrain `json:"grainBill"`
	BatchSize float64 `json:"batchSize"`
}

type SrmResponse struct {
	Srm float64 `json:"srm"`
	Ebc float64 `json:"ebc"`
	HexColor string `json:"hexColor"`
}

type RecipeGrain struct {
	Name string `json:"name"`
	Amount float64 `json:"amount"`
}

type SrmHex struct {
	SrmKey int `json:"srmKey"`
	HexValue string `json:"hexValue"`
}

func grainColorList(fermentables map[string]Fermentable, recipeGrains []RecipeGrain) []float64 {
	var grainColors []float64
	for _, r := range recipeGrains {
		fermentable := fermentables[strings.ToLower(r.Name)]
		grainColors = append(grainColors, fermentable.DegreesLovibond)
	}
	return grainColors
}

func mcu(grainColor float64, amountLbs float64, batchSize float64) float64 {
	return (grainColor * amountLbs) / batchSize
}

func srmColor(grainColors []float64, grainAmounts []float64, batchSize float64) float64 {
	var totalMcu float64
	for i, c := range grainColors {
		totalMcu += mcu(c, grainAmounts[i], batchSize)
	}
	return math.Pow(totalMcu, 0.6859) * 1.4922
}

func ebcColor(srm float64) float64 {
	return srm * 1.97
}

func srmHexValue(srm float64) string {
	hexKey := int(srm)
	hexValues := retrieveSrmHexValues()
	if srmHex, found := hexValues[hexKey]; found {
		return srmHex.HexValue
	}

	return ""
}

func calculateSrm(request SrmRequest) SrmResponse {
	fermentables := retrieveFermentablesNameKey()

	var grainAmounts []float64
	for _, r := range request.GrainBill {
		grainAmounts = append(grainAmounts, r.Amount)
	}

	grainColors := grainColorList(fermentables, request.GrainBill)

	srm := srmColor(grainColors, grainAmounts, request.BatchSize)
	ebc := ebcColor(srm)
	hexColor := srmHexValue(srm)

	return SrmResponse{
		Srm:      srm,
		Ebc:      ebc,
		HexColor: hexColor,
	}
}