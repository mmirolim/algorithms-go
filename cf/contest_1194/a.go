package main

import "fmt"

// http://codeforces.com/contest/1194/problem/A
func main() {
	var t int
	fmt.Scanf("%d\n", &t)
	for i := 0; i < t; i++ {
		var n, x int
		fmt.Scanf("%d %d\n", &n, &x)
		fmt.Println(x + x)
	}
}
