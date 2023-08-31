package main

import "golang.org/x/exp/slices"

func ContainsNumberWithSlice(s []int, n int) bool {
	contain := slices.Contains(s, n)
	//println("Contains value = ", contain)
	return contain
}
