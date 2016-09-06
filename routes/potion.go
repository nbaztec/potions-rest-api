package routes

import (
	"math"
)

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func mixPotionItems(items ...Item) float64 {
	t := 0.0
	x := 0.0
	for _, item := range items {
		c := float64(math.Max(float64(item.Toxicity), 0.0))
		t += 2 * c * c
		x += 5 * math.Abs(math.Log(math.Max(1, c)))
	}
	return toFixed(math.Sqrt(t)+x, 2)
}
