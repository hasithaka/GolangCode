package main

import "fmt"

type hotdog int

var x hotdog
var y int

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 10
	fmt.Println(x)
	fmt.Println(y)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n",y)

}
