package main

import "fmt"

var age = 35
var a, b = 35, "petr"

func test(chis int, st int) int {
	temp := chis

	for i := 1; i < st; i++ {
		chis = chis * temp
	}
	return chis
}

func main() {
	fmt.Println("Hello world")
	fmt.Println("Second line")
	fmt.Printf("Привет %s %d %.3f\n", "Пётр", age, 2.5745)
	fmt.Printf("Привет %T \n", age)
	fmt.Printf("%d %s\n", a, b)
	fmt.Printf("%d\n", test(2, 3))
}
