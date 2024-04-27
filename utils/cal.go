package utils

import "fmt"

func sayHello(name string) {

	fmt.Println("Hi, ", name)

	AddName("kk")
}

func AddName(name string) {

	fmt.Println("Hi, ", name)
}
