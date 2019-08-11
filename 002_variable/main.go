package main

import "fmt"

var y = 10

func main() {

	x := 5
	fmt.Println(x)
	foo()
	y = 100
	foo()

	var st = "Hello"
	fmt.Println(st)

	st = st + ` something better " world "`
	fmt.Println(st)
	fmt.Printf("%T\n", st)

	const version string = "1.2.0"

	fmt.Println(version)
	//version = "1.2.2"  --> unable to assign value

	fmt.Println(version)

	fmt.Print("Please enter a number:")
	var input float32
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println("Multiplication of you number is ", output)
	fmt.Printf("%T\n", &input)
	fmt.Printf("%T\n", input)

}

func foo() {

	fmt.Println(y)
}
