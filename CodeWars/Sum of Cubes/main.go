package main

import "fmt"

func SumCubes(n int) int {
	var res int
	for i := 1; i <= n; i++ {
		res += i * i * i
	}
	return res
}

func main() {
	fmt.Println(SumCubes(0))
	fmt.Println(SumCubes(1))
	fmt.Println(SumCubes(3))

	SumCubes(0)
	SumCubes(1)
	SumCubes(3)
}
