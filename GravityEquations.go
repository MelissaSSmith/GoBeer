package main

func gravityPoints(specificGravity float32) float32 {
	return (specificGravity - 1.0) * 1000.0
}

func specificGravity(gravityPoints float32) float32 {
	return (gravityPoints / 1000.0) + 1.0
}