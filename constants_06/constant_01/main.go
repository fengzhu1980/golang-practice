package main

import "fmt"

const p string = "first const"

func main() {
	const q = 42

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)
}
