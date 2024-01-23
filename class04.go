package main

import (
	"fmt"
)

func main() {
	fixedSizeOfFiveColors()
	fixedSizeOfFiveIntegers()
	fixedSizeofFiveNumbers()
	checkAnArray()
	inputExample()
 
}

func fixedSizeofFiveNumbers(){
	//Fixed-size array of 5 numbers

	numbers := [5]float64{76.5, 88.2, 90.0, 65.4, 82.3}

	//summing each element individually
 
	sum := numbers[0] + numbers[1] + numbers[2] + numbers[3] + numbers[4]
 
	//Calculate average
 
	average := sum / 5
 
	fmt.Printf("Numbers: %v\n", numbers)
	fmt.Printf("Sum: %.2f\n", sum)
	fmt.Printf("Average: %.2f\n", average)
	fmt.Printf("/////////////////////////////////////////\n\n")
}

func fixedSizeOfFiveColors(){
	//Fixed size array of 5 colors

	colors := [5]string{"Red", "Blue", "Green", "Yellow", "Purple"}

	//Counting the colors
	count := len(colors)
 
	fmt.Printf("Colors: %v\n", colors)
	fmt.Printf("Number of colors: %d\n", count)
	fmt.Printf("/////////////////////////////////////////\n\n")
 

}

func fixedSizeOfFiveIntegers(){
	 //Fixed-size array of 5 integers

	 numbers := [5]int{76, 88, 90, 65, 82}

	 //Summing each element individually
	
	 sum := numbers[0] + numbers[1] + numbers[2] + numbers[3] + numbers[4]
	
	 //Calculating average
	
	 average := float64(sum) / 5
	
	 fmt.Printf("Numbers: %v\n", numbers)
	 fmt.Printf("Sum: %d\n", sum)
	 fmt.Printf("Average: %.2f\n", average)
	 fmt.Printf("/////////////////////////////////////////\n\n")	
}

func checkAnArray(){
/*Create an array - string of ice cream flavours, count the length, print the number and array.
Bonus: Check if the array contains the value for “vanilla”*/
	flavours := [5]string{"strawberry", "chocolate", "vanilla", "mint", "rocky road"}
	//checkIfContains := "vanila"

	count  := len(flavours)
	fmt.Printf("Count: %v\n", count)
	fmt.Println(flavours)
	fmt.Printf("/////////////////////////////////////////\n\n")
}


func inputExample() {

 //Prices of bakery items

 priceBread := 2.00
 priceCake := 15.00
 priceCookiesPerDozen := 5.00

 //Sales tax rate

 salesTaxRate := 0.08

 //Quantities ordered

 var quantityBread, quantityCake, quantityCookiesDozen int

 //User inputs for quantities

 fmt.Print("Enter quantity of bread: ")

 fmt.Scanln(&quantityBread)
 
 fmt.Print("Enter quantity of cake: ")

 fmt.Scanln(&quantityCake)

 fmt.Print("Enter quantity of cookies:")

 fmt.Scanln(&quantityCookiesDozen)

 //Calculate the total cost before tax

 totalCostBeforeTax := (priceBread * float64(quantityBread)) + (priceCake * float64(quantityCake)) + (priceCookiesPerDozen * float64(quantityCookiesDozen))

 //Calculate the total cost after tax

 totalCostAfterTax := totalCostBeforeTax + (totalCostBeforeTax * salesTaxRate)

 //Display the results

 fmt.Printf("Total cost before tax: $%.2f\n", totalCostBeforeTax)
 fmt.Printf("Total cost after tax: $%.2f\n", totalCostAfterTax)

}