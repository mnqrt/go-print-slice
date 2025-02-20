package main

import "github.com/mnqrt/go-print-slice/printslice"

func main() {
	a := make([][]int, 2)
	a[0] = make([]int, 2)
	a[1] = make([]int, 2)
	a[0][0] = 1
	a[0][1] = 2
	a[1][0] = 90909
	a[1][1] = 9
	printslice.PrintSlice(a[1])
	printslice.Print2DSlice(a)
}