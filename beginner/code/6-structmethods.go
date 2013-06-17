package main

import "fmt"

type MyCar struct {
	speed int
}

func (m *MyCar) acc() {
	m.speed = m.speed + 10
}

func main() {
	m := MyCar{}
	fmt.Println(m)

	m.acc()
	fmt.Println(m)
}

// 1. dot notation to reach functionality
// 2. pointers automatically resolved in many cases.  You didn't have to make it &m to make the call.  (Pointers not always automatically resolved)
// 3. Functions are associated with a type.  They are not IN a type like in Java.
