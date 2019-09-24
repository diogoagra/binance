package main

import (
	"math"
)

const feeBINANCE = 0.10

func orderCalc(kind string, quantity, rate float64) (total float64, totalwithoutfee float64) {

	if kind == "buy" {
		total = (quantity / rate)

	}
	if kind == "sell" {
		total = (rate * quantity)
	}

	fee := (total / 100) * feeBINANCE

	total = ToFixed(total-fee, 8)
	totalwithoutfee = ToFixed(total, 8)

	return
}

// Round -
func Round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

// ToFixed -
func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(Round(num*output)) / output
}
