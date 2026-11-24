package main

import "fmt"

// CanMergeTriplets reports whether the target triplet can be produced by
// repeatedly merging chosen triplets with element-wise max. A triplet is
// usable iff it nowhere exceeds the target; the target is reachable iff the
// usable triplets jointly supply all three exact coordinates.
func CanMergeTriplets(triplets [][3]int, target [3]int) bool {
	var hit [3]bool
	for _, t := range triplets {
		if t[0] > target[0] || t[1] > target[1] || t[2] > target[2] {
			continue // would overshoot some coordinate forever: unusable
		}
		for i := 0; i < 3; i++ {
			if t[i] == target[i] {
				hit[i] = true
			}
		}
	}
	return hit[0] && hit[1] && hit[2]
}

func main() {
	fmt.Println(CanMergeTriplets([][3]int{{2, 5, 3}, {1, 8, 4}, {1, 7, 5}}, [3]int{2, 7, 5}))
	fmt.Println(CanMergeTriplets([][3]int{{3, 4, 5}, {4, 5, 6}}, [3]int{3, 2, 5}))
}
