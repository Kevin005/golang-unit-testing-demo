package main

import "fmt"

type HelloWorld struct{}

func (_ *HelloWorld) PrintHello() {
	fmt.Println("hello")
}

func (_ *HelloWorld) PrintWorld() {
	fmt.Println("world")
}
