package main

import "github.com/thoas/go-funk"

func ContainsNumberWithFunk(s []int, n int) bool {
	contain := funk.Contains(s, n)
	//println("Contains value = ", contain)
	return contain
}
