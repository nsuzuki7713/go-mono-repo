package main

import (
	"fmt"

	"hello-world/calc1/sum1"

	"hello-world/mypkg"
)

func main() {
	fmt.Println("Hello World")

	mypkg.PrintBar()
	
	sum1.Summarize([]int{1, 2, 3, 4, 5})
}