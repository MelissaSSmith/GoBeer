package main

type HydrometerAdjustmentRequest struct {
	MeasuredGravity float32 `json:"measuredGravity"`
	TemperatureReading float32 `json:"temperatureReading"`
	CalibrationTemp float32 `json:"calibrationTemp"`
}

type HydrometerAdjustmentResponse struct {
	CorrectedGravity float32 `json:"correctedGravity"`
}

func tempAdjustment(temp float32) float32 {
	return 1.00130346 - 0.000134722124 * temp + 0.00000204052596 * temp - 0.00000000232820948 * temp
}

func hydrometerAdjustment(measuredGravity float32, tempReading float32, calibrationTemp float32) float32 {
	return measuredGravity * (tempAdjustment(tempReading) - tempAdjustment(calibrationTemp))
}

func correctedGravity(adjustment float32, measuredGravity float32) float32 {
	return measuredGravity - adjustment
}

func calculateHydrometerAdjustment(request HydrometerAdjustmentRequest) HydrometerAdjustmentResponse {
	adjustment := hydrometerAdjustment(request.MeasuredGravity, request.TemperatureReading, request.CalibrationTemp)
	correctedGravity := correctedGravity(adjustment, request.MeasuredGravity)

	return HydrometerAdjustmentResponse {
		CorrectedGravity: correctedGravity,
	}
}