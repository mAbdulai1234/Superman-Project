package main
import "fmt"

/*
Create a program that asks the user for unlimited product dta (name, price, stock)
the programe should have amenu with this options

1- add product
2- show product by name
3 remove product
4- edit product
5- list all products
6-Exit
*/
func main(){
	names := []string{}
	prices := []float32{}
	stocks := [] int{}

	choice :=0

	for true {
		fmt.Println("1 - Add product")
		fmt.Println("1 - Show product by name")
		fmt.Println("1 - Remove product")
		fmt.Println("1 - Edit product")
		fmt.Println("1 - List all product")
		fmt.Println("1 - exit product")

		fmt.Println()

		fmt.Println("Please enter an option")
		fmt.Scan(&choice)

		if choice ==1 {

			fmt.Println("Enter the name of the new product")
			var name string
			fmt.Scan(&name)

			fmt.Println("Enter the name of the new product")
			var price float32
			fmt.Scan(&price)

			fmt.Println("Enter the name of the new product")
			var stock int
			fmt.Scan(&stock)

			names = append(names, name)
			prices = append(prices, price)
			stocks = append(stocks, stock)

			

		} else if choice ==2{

		} else if choice ==3{

		} else if choice ==4{

		} else if choice ==5{

		} else if choice ==6{
			break
		}
		


	}
}

