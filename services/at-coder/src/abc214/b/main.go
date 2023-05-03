package main

import "fmt"

func main() {
	var s, t int

	fmt.Scanf("%d %d", &s, &t)

	count := 0
	for i := 0; i <= s; i++ {
		for j := 0; j <= s-i; j++ {
			for k := 0; k <= s-i-j; k++ {
				if i+j+k <= s && i*j*k <= t {
					count++
				}
			}
		}
	}

	fmt.Println(count)
}
