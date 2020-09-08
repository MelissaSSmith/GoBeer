package main

type HydrometerAdjustmentRequest struct {
	MeasuredGravity float64 `json:"measured_gravity"`
	TemperatureReading float64 `json:"temperature_reading"`
	CalibrationTemp float64 `json:"calibration_temp"`
}

type HydrometerAdjustmentResponse struct {
	CorrectedGravity float64 `json:"corrected_gravity"`
}

func tempAdjustment(temp float64) float64 {
	return 1.00130346 - 0.000134722124 * temp + 0.00000204052596 * temp - 0.00000000232820948 * temp
}

func hydrometerAdjustment(measuredGravity float64, tempReading float64, calibrationTemp float64) float64 {
	return measuredGravity * (tempAdjustment(tempReading) - tempAdjustment(calibrationTemp))
}

func correctedGravity(adjustment float64, measuredGravity float64) float64 {
	return measuredGravity - adjustment
}

func calculateHydrometerAdjustment(request HydrometerAdjustmentRequest) HydrometerAdjustmentResponse {
	adjustment := hydrometerAdjustment(request.MeasuredGravity, request.TemperatureReading, request.CalibrationTemp)
	correctedGravity := correctedGravity(adjustment, request.MeasuredGravity)

	return HydrometerAdjustmentResponse {
		CorrectedGravity: correctedGravity,
	}
}