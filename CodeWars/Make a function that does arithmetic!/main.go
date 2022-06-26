package main

import "fmt"

func Arithmetic(a int, b int, operator string) int {
	var res int
	switch {
	case operator == "add":
		res = a + b
	case operator == "subtract":
		res = a - b
	case operator == "multiply":
		res = a * b
	case operator == "divide":
		res = a / b
	}
	return res
}

func main() {
	fmt.Println(Arithmetic(8, 2, "add"))
	fmt.Println(Arithmetic(8, 2, "subtract"))
	fmt.Println(Arithmetic(8, 2, "multiply"))
	fmt.Println(Arithmetic(8, 2, "divide"))

	//Arithmetic(8, 2, "add")
	//Arithmetic(8, 2, "subtract")
	//Arithmetic(8, 2, "multiply")
	//Arithmetic(8, 2, "divide")

}
