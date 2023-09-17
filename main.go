package main

import (
	"fmt"
)

// type owner struct {
// 	name string
// 	age  uint8
// }

// type mascota struct {
// 	name string
// 	owner
// }

// func main() {
// 	perro := mascota{
// 		owner: owner{
// 			name: "Noah",
// 			age:  24,
// 		},
// 		name: "Pochi",
// 	}

// 	fmt.Println(perro.name)
// }

func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("world"))
}
