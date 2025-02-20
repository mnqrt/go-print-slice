package main

import "github.com/mnqrt/go-print-slice/printslice"

func main() {
	a := make([]int, 3)
	a[0], a[1], a[2] = 2, 4, 1
	printslice.PrintSlice(a)
}