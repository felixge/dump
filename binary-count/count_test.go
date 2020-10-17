package count

import "testing"

func TestBinaryCount(t *testing.T) {
	tests := []struct {
		N    int
		Want int
	}{
		{9, 2},
		{529, 4},
		{20, 1},
		{15, 0},
		{32, 0},
		{1041, 5},
	}
	for _, test := range tests {
		got := Count(test.N)
		if got != test.Want {
			t.Fatalf("n=%d got=%d want=%d", test.N, got, test.Want)
		}
	}
}
