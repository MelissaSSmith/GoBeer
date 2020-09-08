package main

import (
	"math"
	"strings"
)

type IbuRequest struct {
	TargetOriginalGravity float64 `json:"target_original_gravity"`
	BoilSize float64 `json:"boil_size"`
	BatchSize float64 `json:"batch_size"`
	Hops []RecipeHop `json:"hops"`
}

type IbuResponse struct {
	EstimatedBoilGravity float64 `json:"estimated_boil_gravity"`
	TotalIbu float64 `json:"total_ibu"`
	HopIbus []HopIbu `json:"hop_ibus"`
}

type RecipeHop struct {
	Name string `json:"name"`
	Ounces float64 `json:"ounces"`
	AlphaAcids float64 `json:"alpha_acids"`
	BoilTime float64 `json:"boil_time"`
	Type string `json:"type"`
}

type HopIbu struct {
	Utilization float64 `json:"utilization"`
	Ibus float64 `json:"ibus"`
}

func adjustForPelletHops(hopType string, ibu float64) float64 {
	if strings.ToLower(hopType) == "pellet" {
		return ibu + (ibu * 0.10)
	}
	return ibu
}

func hopUtilization(time float64, boilGravity float64) float64 {
	gravityFunc := 1.65 * math.Pow(0.000125, boilGravity - 1.0)
	timeFunc := (1.0 - math.Pow(math.E, -0.04 * time)) / 4.15
	return gravityFunc * timeFunc
}

func boilGravity(targetedOriginalGravity float64, boilSize float64, batchSize float64) float64 {
	return ((batchSize * math.Mod(targetedOriginalGravity, 1.0)) / boilSize) + 1.0
}

func calculateHopUtilizations(recipeHops []RecipeHop, boilGravity float64) map[string]float64 {
	hopUtilMap := make(map[string] float64)
	for _, h := range recipeHops {
		hopUtilMap[h.Name] = hopUtilization(h.BoilTime, boilGravity)
	}
	return hopUtilMap
}

func calculateHopIbus(hopUtilizations map[string]float64, recipeHops []RecipeHop, finalVolume float64) ([]HopIbu, float64) {
	var hopIbus []HopIbu
	var totalIbu float64
	for _, r := range recipeHops {
		ibu := r.Ounces * r.AlphaAcids * hopUtilizations[r.Name] * (75 / finalVolume)
		ibu = adjustForPelletHops(r.Type, ibu)
		totalIbu += ibu
		hopIbu := HopIbu{
			Utilization: hopUtilizations[r.Name],
			Ibus:        ibu,
		}
		hopIbus = append(hopIbus, hopIbu)
	}
	return hopIbus, totalIbu
}

func calculateIbu(request IbuRequest) IbuResponse {

	boilGravity := boilGravity(request.TargetOriginalGravity, request.BoilSize, request.BatchSize)
	hopUtils := calculateHopUtilizations(request.Hops, boilGravity)
	hopIbus, totalIbu := calculateHopIbus(hopUtils, request.Hops, request.BatchSize)

	return IbuResponse{
		EstimatedBoilGravity: boilGravity,
		TotalIbu:             totalIbu,
		HopIbus:			  hopIbus,
	}
}