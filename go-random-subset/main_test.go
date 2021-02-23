package main

import (
	"fmt"
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

	in := makeSlice(n)
	inMean := mean(in)

	for _, a := range algorithms {
		t.Run(a.Name, func(t *testing.T) {
			for k := 1; k < n*2; k += n / 100 {
				var (
					outMeanSum float64
					minErr     float64
					iterations = n * 10
				)
				for i := 0; i < iterations; i++ {
					dupes := map[int]struct{}{}
					out := a.Fn(k, in)
					for _, v := range out {
						if _, dupe := dupes[v]; dupe {
							t.Fatal("dupe detected")
						}
						dupes[v] = struct{}{}
					}

					if (k <= n && len(out) != k) || (k > n && len(out) != n) {
						t.Fatal("bad output length")
					}

					outMeanSum += mean(out)
					outMean := outMeanSum / float64(i+1)
					outMeanError := math.Abs(inMean-outMean) / inMean
					if outMeanError < tolerance {
						//t.Logf("k=%d i=%d", k, i)
						break
					} else if i+1 == iterations {
						t.Fatalf("sample mean did not converge on population mean: k=%d n=%d iterations=%d minErr=%f", k, n, iterations, minErr)
					} else if outMeanError < minErr || minErr == 0 {
						minErr = outMeanError
					}
				}
			}
		})
	}
}

func BenchmarkAlgorithms(b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	n := 100000
	for _, a := range algorithms {
		for p := float64(1); p < 100; p += 9.8 {
			k := int(float64(n) * (p / 100))

			name := fmt.Sprintf("%s-%d", a.Name, int(math.Round(p)))
			b.Run(name, func(b *testing.B) {
				input := makeSlice(n)
				for i := 0; i < b.N; i++ {
					a.Fn(k, input)
				}
			})
		}
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
