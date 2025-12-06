package main

import "fmt"

func constant() {
	// const name = "Andi Lauw"
	// const ttl = "Lampung"

	// fmt.Println(name)


	// Multiple Constant 
	const (
		name = "Andi Lauw"
		ttl = "Lampung"
	)
	fmt.Println(name + ttl)
}