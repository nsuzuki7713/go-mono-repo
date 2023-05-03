// https://atcoder.jp/contests/abc165/tasks/abc165_a
package main

import "fmt"

func main() {
	var k, a, b int
	fmt.Scan(&k)
	fmt.Scan(&a, &b)

	for i := a; i <= b; i++ {
		if i%k == 0 {
			fmt.Println("OK")
			return
		}
	}

	fmt.Println("NG")
}
