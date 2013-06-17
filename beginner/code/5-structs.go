package main

import "fmt"

type MyCar struct {
	color    string
	maxSpeed int
}

func main() {
	m := MyCar{}
	fmt.Println(m) //{ 0}

	m = MyCar{"red", 100}
	fmt.Println(m) //{red, 100}

	m.color = "blue"
	m.maxSpeed = 150
	fmt.Println(m) //{blue, 150}
	fmt.Println("color is:", m.color)
	//color is: blue

	m = MyCar{maxSpeed: 150, color: "green"}
	fmt.Println(m) //{green, 150}
}
