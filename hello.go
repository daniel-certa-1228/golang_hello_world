package main

import (
	"fmt"
	"math"
	"reflect"
	"unsafe"
)

func main() {
	//Hello, world.
	fmt.Printf("hello, world\n")
	//Variables
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
	//Types
	bool_1 := true
	bool_2 := false
	fmt.Println("bool_1:", bool_1, "bool_2:", bool_2)
	bool_3 := bool_1 && bool_2
	fmt.Println("bool_3:", bool_3)
	bool_4 := bool_1 || bool_2
	fmt.Println("bool_4:", bool_4)

	var int_1 int = 100
	int_2 := 200
	fmt.Println("int_1:", int_1, "int_2", int_2)
	fmt.Printf("type of int_1 is %T, size of int_1 is %d\n", int_1, unsafe.Sizeof(int_1))
	fmt.Printf("type of int_2 is %T, size of int_2 is %d\n", int_2, unsafe.Sizeof(int_2))

	var first string = "Dan"
	last := "Certa"
	fmt.Println(first + " " + last)

	var i int = 100
	var j float64 = 25.99
	fmt.Println(float64(i) + j)

	const con1 = 12
	const con2 = "Con 2"

	var name1 = "Charlie"
	type myString string
	var customeName myString = "Charlie"

	fmt.Println(name1, customeName)
}
