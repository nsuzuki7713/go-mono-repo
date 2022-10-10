package main

import (
	"fmt"

	"hello-world/calc1"
	"hello-world/mypkg"

	"github.com/geek-line/sum"
)

func main() {
	fmt.Println("Hello World")

	mypkg.PrintBar()

	fmt.Println(calc1.Summarize([]int{1, 2, 3, 4, 5}))

	fmt.Println(sum.Sum(2, 3))
}
