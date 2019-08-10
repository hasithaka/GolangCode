package main

import "fmt"

func main() {
	fmt.Println("Hello, This is from main function")
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println("This is from for loop")
		}
	}
	foo()
}

func foo() {
	n, err := fmt.Println("This is from function foo")
	fmt.Println(n)
	fmt.Println(err)

}
