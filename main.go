package main

import "fmt"

func main() {
	var (
		bubblegum     = 202
		toffee        = 118
		iceCream      = 2250
		milkChocolate = 1680
		doughunt      = 1075
		pancake       = 80
	)

	var staffExpenses, otherExpenses int

	income := bubblegum + toffee + iceCream + milkChocolate + doughunt + pancake
	fmt.Println("Earned amount:")
	fmt.Printf("Bubblegum: $%d\n", bubblegum)
	fmt.Printf("Toffee: $%d\n", toffee)
	fmt.Printf("Ice cream: $%d\n", iceCream)
	fmt.Printf("Milk chocolate: $%d\n", milkChocolate)
	fmt.Printf("Doughnut: $%d\n", doughunt)
	fmt.Printf("Pancake: $%d\n", pancake)
	fmt.Printf("\nIncome: $%d\n", income)

	fmt.Println("Staff expenses:")
	fmt.Scan(&staffExpenses)
	fmt.Println("Other expenses:")
	fmt.Scan(&otherExpenses)

	netIncome := income - staffExpenses - otherExpenses
	fmt.Printf("Net income: $%d\n", netIncome)
}
