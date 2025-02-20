package main

import "fmt"

func main() {
	fmt.Println("HALOO")
	PrintFormat("BEBEK")
}

func PrintFormat(s string) {
	fmt.Println("---", s, "---")
}