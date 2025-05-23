package main

import "fmt"

func get() string {
	return "phani"
}

func main() {
	a := get()
	fmt.Println("the address was ",&a)
	b := &a
	*b = "phanidhar" // in the same memory address we are storing a new value by overriding the old one
	
}