package main

import (
	"math/rand"
	"time"
)

func generateSample(sampleSize int, weights []float64, upperThreshold float64, lowerThreshold float64) ([][]float64, []int) {
	samples := make([][]float64, 0, sampleSize)
	labels := make([]int, 0, sampleSize)
	rand.Seed(time.Now().UnixNano())
	for len(samples) < sampleSize {
		newSample := make([]float64, len(weights))
		var sum float64
		for j, w := range weights {
			r := rand.Float64()*2 - 1
			sum = sum + r*w
			newSample[j] = r
		}
		if sum > upperThreshold {
			samples = append(samples, newSample)
			labels = append(labels, 1)
		} else if sum < lowerThreshold {
			samples = append(samples, newSample)
			labels = append(labels, -1)
		}
	}
	return samples, labels
}
