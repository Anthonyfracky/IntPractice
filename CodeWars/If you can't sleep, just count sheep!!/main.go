package main

import (
	"strconv"
)

func countSheep(num int) string {
	var res string
	if num == 0 {
		res = ""
		return res
	} else {
		for i := 1; i <= num; i++ {
			tempInt := i
			tempStr := strconv.Itoa(tempInt)
			res = res + tempStr + " sheep..."
		}
		return res
	}
}

func main() {
	countSheep(0)
	countSheep(1)
	countSheep(3)
	countSheep(5)

	//fmt.Println(countSheep(0))
	//fmt.Println(countSheep(1))
	//fmt.Println(countSheep(3))
	//fmt.Println(countSheep(5))
}
