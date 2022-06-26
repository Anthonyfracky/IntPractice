package main

import "fmt"

func main() {
	arr := []int{4, 1, 4, -4, 6, 3, 8, 8}
	temper := []int{}
	sorter := []int{}
	q := 0

	for i := 0; i < len(arr); i++ {
		for j := 0; j < 20; j++ {
			temper = append(temper, arr[i])
			if arr[i] != temper[q] {
				sorter = append(sorter, arr[i])
				fmt.Println("+")
			}
			q++
		}
	}

	result := sorter[:]
	fmt.Println(result)
	fmt.Println(sorter)
}
