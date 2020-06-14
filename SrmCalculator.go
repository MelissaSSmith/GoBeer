package main

import (
	"math"
	"sort"
)

type SrmRequest struct {
	GrainBill []RecipeGrain `json:"grainBill"`
	BatchSize float32 `json:"batchSize"`
}

type SrmResponse struct {
	Srm float32 `json:"srm"`
	Ebc float32 `json:"ebc"`
	HexColor string `json:"hexColor"`
}

type RecipeGrain struct {
	Name string `json:"name"`
	Amount float32 `json:"amount"`
}

type SrmHex struct {
	SrmKey int `json:"srmKey"`
	HexValue string `json:"hexValue"`
}

func grainColorList(fermentables []Fermentable, recipeGrains []RecipeGrain) []float32 {
	var grainColors []float32
	for _, r := range recipeGrains {
		sort.SliceStable(fermentables, func(i, _ int) bool { return fermentables[i].Name == r.Name })
		grainColors = append(grainColors, fermentables[0].DegreesLovibond)
	}
	return grainColors
}

func mcu(grainColor float32, amountLbs float32, batchSize float32) float32 {
	return (grainColor * amountLbs) / batchSize
}

func srmColor(grainColors []float32, grainAmounts []float32, batchSize float32) float32 {
	var totalMcu float32
	for i, c := range grainColors {
		totalMcu += mcu(c, grainAmounts[i], batchSize)
	}
	return float32(math.Pow(float64(totalMcu), 0.6859) * 1.4922)
}

func ebcColor(srm float32) float32 {
	return srm * 1.97
}

func srmHexValue(srm float32) string {
	hexKey := int(srm)
	hexValues := retrieveSrmHexValues()
	sort.SliceStable(hexValues, func(i, _ int) bool { return hexValues[i].SrmKey == hexKey })

	return hexValues[0].HexValue
}

func calculateSrm(request SrmRequest) SrmResponse {
	fermentables := retrieveFermentables()

	var grainAmounts []float32
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