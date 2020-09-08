package main


type AbvRequest struct {
	OriginalGravity	float64	`json:"original_gravity"`
	FinalGravity float64 `json:"final_gravity"`
}

type AbvResponse struct {
	StandardAbv float64 `json:"standard_abv"`
	AlternateAbv float64 `json:"alternate_abv"`
	TotalCalories float64 `json:"total_calories"`
}

func calculateStandardAbv(originalGravity float64, finalGravity float64) float64 {
	return (originalGravity - finalGravity) * 131.25
}

func calculateAlternateAbv(originalGravity float64, finalGravity float64) float64 {
	return (76.08 * (originalGravity - finalGravity) / (1.775 - originalGravity)) * (finalGravity / 0.794)
}

func calculateTotalCalories(originalGravity float64, finalGravity float64) float64 {
	caloriesFromCarbs := 3500.0 * finalGravity * ((0.1808 * originalGravity) + (0.819 * finalGravity) - 1.0004)
	caloriesFromAlcohol := 1881.22 * finalGravity * (originalGravity - finalGravity)/(1.775 - originalGravity)

	return caloriesFromCarbs + caloriesFromAlcohol
}

func calculateAbv(request AbvRequest) AbvResponse {
	standardAbv := calculateStandardAbv(request.OriginalGravity, request.FinalGravity)
	alternateAbv := calculateAlternateAbv(request.OriginalGravity, request.FinalGravity)
	totalCalories := calculateTotalCalories(request.OriginalGravity, request.FinalGravity)

	return AbvResponse {
		StandardAbv:   standardAbv,
		AlternateAbv:  alternateAbv,
		TotalCalories: totalCalories,
	}
}