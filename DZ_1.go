package main

import "fmt"

func main() {
	var (
		applePrice float64 = 5.99
		pearPrice  int     = 7
		budget     float64 = 23
	)

	fmt.Println("\nОдне яблуко коштує", applePrice, "грн. Ціна однієї груші -", pearPrice, "грн.")
	fmt.Println("Ми маємо ", budget, " грн.", "\n")

	fmt.Println("1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?")
	var res1 float64 = (9 * applePrice) + (8 * float64(pearPrice))
	fmt.Println("Щоб купити 9 яблук та 8 груш потрібно", res1, "\n")

	fmt.Println("2. Скільки груш ми можемо купити?")
	var res2 int = int(budget) / int(pearPrice)
	fmt.Println("Ми можемо купити", res2, "груш.", "\n")

	fmt.Println("3. Скільки яблук ми можемо купити?")
	var res3 int = int(budget) / int(applePrice)
	fmt.Println("Ми можемо купити", res3, "яблук.", "\n")

	fmt.Println("4. Чи ми можемо купити 2 груші та 2 яблука?")
	var summ = (2 * applePrice) + (2 * float64(pearPrice))
	if summ < budget {
		fmt.Println("Так")
	} else {
		fmt.Println("Ні")
	}

}
