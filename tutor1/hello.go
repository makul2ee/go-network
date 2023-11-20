package main

import (
	"fmt"
)
// func Init() {
// 	fmt.Println("x= ", x)
// }
func main() {

	// var t bool = true
	// var f bool 

	// names := "MaKul2ee"
	var age int = 40 
	var favNum float64 = 1.6180339

	var complex complex128 = 5+5i

	var r rune = 10

	fmt.Println("age is", age, "favNum is",favNum)
	fmt.Println("complex is", complex)
	fmt.Println("rune is", r)

	// Init()

	var  myName string = "อิอิๆๆ"
	fmt.Println(myName + "is a robot")
	fmt.Println(myName[3])
	fmt.Println("length of myName is" , len(myName))

	var arry5[5] float64
	arry5[0] = 98.7
	arry5[1] = 93.2
	arry5[2] = 99.7
	arry5[3] = 77.2
	arry5[4] = 83.7
	fmt.Println(arry5)
	fmt.Println("lenght of array5 is" , len(arry5))
	fmt.Println("arry5[3] is " ,arry5[3])

	arry3 :=[3]float64 {98,93,99}
	fmt.Println(arry3)


	var s []float64 = arry5[0:5]
	fmt.Println((s))

	type Person struct {
		name string
		age int
	}

	var p Person
	p.name = "Makul2ee"
	p.age = 21
	fmt.Println(p)


	var x int = 5
	var xPtr *int = &x
	fmt.Println(xPtr)
}