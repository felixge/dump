package main

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

// TestCorrectness uses the central limit theorem to check that all subset
// algorithms seem to produce reasonably random values. Somebody with better
// stats skills than me should probably verify that this is a reasonable test.
func TestCorrectness(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	n := 1000
	tolerance := 0.01

	original := makeSlice(n)
	originalMean := mean(original)

	for _, a := range algorithms {
		t.Run(a.Name, func(t *testing.T) {
			for k := 1; k < n; k += n / 100 {
				in := make([]int, len(original))
				copy(in, original)

				var (
					sampleMeanSum float64
					sampleCount   int
					minErr        float64
				)
				for i := 0; i < n*100; i++ {
					sampleMeanSum += mean(a.Fn(k, in))
					sampleCount++
					samplesMean := sampleMeanSum / float64(sampleCount)
					samplesMeanError := math.Abs(originalMean-samplesMean) / originalMean
					if samplesMeanError < tolerance {
						return
					} else if samplesMeanError < minErr || minErr == 0 {
						minErr = samplesMeanError
					}
				}
				t.Fatalf("sample mean did not converge on population mean: k=%d n=%d minErr=%f", k, n, minErr)
			}
		})
	}
}

func makeSlice(n int) []int {
	s := make([]int, n)
	for i := 0; i < len(s); i++ {
		s[i] = i
	}
	return s
}

func mean(in []int) float64 {
	var sum int
	for _, v := range in {
		sum += v
	}
	return float64(sum) / float64(len(in))
}
