package main

import "fmt"

func main() {

	fmt.Println("Hello")
	var x int
	type hotdog int

	var b hotdog

	x = 30
	b = 50

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	fmt.Println(b)
	fmt.Printf("%T\n", b)

	 x = int(b)    // b is hotdog type

	fmt.Println(x)
	fmt.Printf("%T\n", x)

}
