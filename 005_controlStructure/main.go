package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, " is even")
		} else {
			fmt.Println(i, " is odd")
		}
	}

	for x := 1; x <= 10; x++ {
		switch x {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		case 5:
			fmt.Println("Five")
		case 6:
			fmt.Println("Six")
		case 7:
			fmt.Println("Seven")
		case 8:
			fmt.Println("Eight")
		case 9:
			fmt.Println("Nine")
		case 10:
			fmt.Println("Ten")
		}
	}

	for y := 1; y <= 100; y++ {
		if y%3 == 0 {
			fmt.Print(y, ", ")
		}

	}
	fmt.Println()

	for z := 1; z <= 100; z++ {
		if z%3 == 0 && z%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if z%3 == 0 {
			fmt.Println("Fizz")
		} else if z%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(z)
		}
	}
}
