package main

import "fmt"

func main() {
	// a := []int16{}
	// fmt.Println("Hello world!")
	a, err := fmt.Print("Hello world!")
	fmt.Println(a)
	if err != nil {
		fmt.Println(err)
	}
}
