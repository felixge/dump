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
	{"fromOffset", fromOffset},
	{"fisherYatesSubset", fisherYatesSubset},
	{"hashDedupe", hashDedupe},
	//{"hashDedupeFlip", hashDedupeFlip},
	{"hybridHashDedupeFisherYates", hybridHashDedupeFisherYates},
	//{"richardExp", richardExp},
	{"algorithmL", algorithmL},
}

func fisherYatesSubset(k int, in []int) []int {
	if k >= len(in) {
		return in
	}

	inCopy := make([]int, len(in))
	copy(inCopy, in)
	rand.Shuffle(len(inCopy), func(i, j int) {
		inCopy[i], inCopy[j] = inCopy[j], inCopy[i]
	})
	return inCopy[0:k]
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

func hashDedupeFlip(k int, in []int) []int {
	if k >= len(in) {
		return in
	}

	m := map[int]struct{}{}
	out := make([]int, 0, k)
	flip := k > len(in)/2

	for len(out) < k {
		i := rand.Intn(len(in))
		if _, dupe := m[i]; dupe {
			continue
		}
		if !flip {
			out = append(out, in[i])
		}
	}
	//if flip {
	//for i, v := range in {
	//if _, ok
	//out = append(out, v)
	//}
	//}
	return out
}

func hashDedupe(k int, in []int) []int {
	if k >= len(in) {
		return in
	}

	picked := map[int]struct{}{}
	out := make([]int, 0, k)
	for len(out) < k {
		i := rand.Intn(len(in))
		if _, dupe := picked[i]; dupe {
			continue
		}
		picked[i] = struct{}{}
		out = append(out, in[i])
	}
	return out
}

func fromOffset(k int, in []int) []int {
	if k >= len(in) {
		return in
	}
	out := make([]int, 0, k)
	offset := rand.Intn(len(in))
	for i := 0; i < k; i++ {
		out = append(out, in[(i+offset)%len(in)])
	}
	return out
}

// https://www.codementor.io/@alexanderswilliams/how-to-efficiently-generate-a-random-subset-150hbz3na4
func hybridHashDedupeFisherYates(k int, in []int) []int {
	if float64(k) < float64(0.63212055882)*float64(len(in)) {
		return hashDedupe(k, in)
	} else {
		return fisherYatesSubset(k, in)
	}
}

// https://en.wikipedia.org/wiki/Reservoir_sampling#An_optimal_algorithm
func algorithmL(k int, s []int) []int {
	if k >= len(s) {
		return s
	}

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
