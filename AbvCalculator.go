package main


type AbvRequest struct {
	OriginalGravity	float32	`json:"originalGravity"`
	FinalGravity float32 `json:"finalGravity"`
}

type AbvResponse struct {
	StandardAbv float32 `json:"standardAbv"`
	AlternateAbv float32 `json:"alternateAbv"`
	TotalCalories float32 `json:"totalCalories"`
}

func calculateStandardAbv(originalGravity float32, finalGravity float32) float32 {
	return (originalGravity - finalGravity) * 131.25
}

func calculateAlternateAbv(originalGravity float32, finalGravity float32) float32 {
	return (76.08 * (originalGravity - finalGravity) / (1.775 - originalGravity)) * (finalGravity / 0.794)
}

func calculateTotalCalories(originalGravity float32, finalGravity float32) float32 {
	caloriesFromCarbs := 3500.0 * finalGravity * ((0.1808 * originalGravity) + (0.819 * finalGravity) - 1.0004)
	caloriesFromAlcohol := 1881.22 * finalGravity * (originalGravity - finalGravity)/(1.775 - originalGravity)

	return caloriesFromCarbs + caloriesFromAlcohol
}

func calculateAbv(request AbvRequest) AbvResponse {
	standardAbv := calculateStandardAbv(request.OriginalGravity, request.FinalGravity)
	alternateAbv := calculateAlternateAbv(request.OriginalGravity, request.FinalGravity)
	totalCalories := calculateTotalCalories(request.OriginalGravity, request.FinalGravity)

	return AbvResponse{
		StandardAbv:   standardAbv,
		AlternateAbv:  alternateAbv,
		TotalCalories: totalCalories,
	}
}