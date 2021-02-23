package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	return nil
}

var algorithms = []struct {
	Name string
	Fn   func(k int, in []int) []int
}{
	{"fisherYatesSubset", fisherYatesSubset},
	//{"richardExp", richardExp},
	{"algorithmL", algorithmL},
}

func fisherYatesSubset(k int, in []int) []int {
	rand.Shuffle(len(in), func(i, j int) {
		in[i], in[j] = in[j], in[i]
	})
	return in[0:k]
}

func richardExp(k int, in []int) []int {
	n := len(in)
	sample := make([]int, k)
	remainingToSampleFrom := n
	next := -1

	for pos := 0; pos < k; pos++ {
		p := float64(k-pos) / float64(remainingToSampleFrom)
		expGap := math.Log(rand.Float64())/math.Log(1-p) + 1
		gap := int(math.Min(float64(n-k+pos-next), expGap))
		next += gap
		sample[pos] = in[next]
		remainingToSampleFrom -= gap
	}
	return sample
}

// https://en.wikipedia.org/wiki/Reservoir_sampling#An_optimal_algorithm
func algorithmL(k int, s []int) []int {
	n := len(s)
	r := make([]int, k)
	i := 0
	for i = 0; i < k; i++ {
		r[i] = s[i]
	}

	w := math.Exp(math.Log(rand.Float64()) / float64(k))
	for i < n {
		i = i + int(math.Floor(math.Log(rand.Float64())/math.Log(1-w))) + 1
		if i < n {
			r[rand.Intn(k)] = s[i]
			w = w * math.Exp(math.Log(rand.Float64())/float64(k))
		}
	}
	return r
}
