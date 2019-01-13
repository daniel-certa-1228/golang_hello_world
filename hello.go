package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Printf("hello, world\n")

	var age int
	fmt.Println("My age is", age)
	age = 44
	fmt.Println("My age is", age)
	age = 100
	fmt.Println("My age is", age)

	var newVar int = 29000
	fmt.Println("newVar =", newVar)

	var inferredVariable_1 = 300
	var inferredVariable_2 = "300"
	var inferredVariable_3 = 300.003

	fmt.Println("Type of inferredVariable_1 is", reflect.TypeOf(inferredVariable_1))
	fmt.Println("Type of inferredVariable_2 is", reflect.TypeOf(inferredVariable_2))
	fmt.Println("Type of inferredVariable_3 is", reflect.TypeOf(inferredVariable_3))

	var width, height int
	fmt.Println("Width is", width, "height is", height)
	width = 50
	height = 100
	fmt.Println("Width is", width, "height is", height)

	name, age := "Dan", 44
	fmt.Println("Name is", name, "age is", age)

	a, b := 145.8, 543.8
	c := math.Min(a, b)
	fmt.Println("minimum value is ", c)
}
