package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if 1 <= n && n <= 125 {
		fmt.Println(4)
	} else if 126 <= n && n <= 211 {
		fmt.Println(6)
	} else {
		fmt.Println(8)
	}
}
