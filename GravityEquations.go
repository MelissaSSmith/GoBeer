package main

func gravityPoints(specificGravity float64) float64 {
	return (specificGravity - 1.0) * 1000.0
}

func specificGravity(gravityPoints float64) float64 {
	return (gravityPoints / 1000.0) + 1.0
}