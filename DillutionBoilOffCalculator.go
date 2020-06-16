package main

import (
	"math"
)

type DilutionBoilOffRequest struct {
	NewVolumeCurrentGravity float64 `json:"new_volume_current_gravity"`
	DesiredGravity float64 `json:"desired_gravity"`
	NewVolumeWortVolume float64 `json:"new_volume_wort_volume"`
	NewGravityWortVolume float64 `json:"new_gravity_wort_volume"`
	NewGravityCurrentGravity float64 `json:"new_gravity_current_gravity"`
	TargetVolume float64 `json:"target_volume"`
}

type DilutionBoilOffResponse struct {
	NewVolume float64 `json:"new_volume"`
	NewVolumeDiff float64 `json:"new_volume_diff"`
	NewGravity float64 `json:"new_gravity"`
	NewGravityDiff float64 `json:"new_gravity_diff"`
}

func newVolume(currentGravity float64, desiredGravity float64, wortVolume float64) float64 {
	currentGravityPoints := gravityPoints(currentGravity)
	desiredGravityPoints := gravityPoints(desiredGravity)

	return (currentGravityPoints * wortVolume) / desiredGravityPoints
}

func newGravity(currentVolume float64, currentGravity float64, targetVolume float64) float64 {
	currentGravityPoints := gravityPoints(currentGravity)
	gravityPoints := (currentGravityPoints * currentVolume) / targetVolume

	return specificGravity(gravityPoints)
}

func calculateBoilOff(request DilutionBoilOffRequest) DilutionBoilOffResponse {
	newVolume := newVolume(request.NewVolumeCurrentGravity, request.DesiredGravity, request.NewVolumeWortVolume)
	newGravity := newGravity(request.NewGravityWortVolume, request.NewGravityCurrentGravity, request.TargetVolume)

	return DilutionBoilOffResponse{
		NewVolume:      newVolume,
		NewVolumeDiff:  math.Abs(newVolume - request.NewVolumeWortVolume),
		NewGravity:     newGravity,
		NewGravityDiff: math.Abs(newGravity - request.NewGravityCurrentGravity),
	}
}
