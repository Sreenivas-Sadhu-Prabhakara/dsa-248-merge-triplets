package main

import "testing"

func TestCanMergeTriplets(t *testing.T) {
	cases := []struct {
		triplets [][3]int
		target   [3]int
		want     bool
	}{
		{[][3]int{{2, 5, 3}, {1, 8, 4}, {1, 7, 5}}, [3]int{2, 7, 5}, true},
		{[][3]int{{3, 4, 5}, {4, 5, 6}}, [3]int{3, 2, 5}, false},
		{[][3]int{{2, 5, 3}, {2, 3, 4}, {1, 2, 5}, {5, 2, 3}}, [3]int{5, 5, 5}, true},
		{[][3]int{{1, 1, 1}}, [3]int{1, 1, 1}, true},
		{[][3]int{{2, 2, 2}}, [3]int{1, 1, 1}, false}, // every triplet overshoots
		{nil, [3]int{1, 2, 3}, false},
	}
	for _, c := range cases {
		if got := CanMergeTriplets(c.triplets, c.target); got != c.want {
			t.Errorf("CanMergeTriplets(%v,%v) = %v, want %v", c.triplets, c.target, got, c.want)
		}
	}
}
