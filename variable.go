package main

import "fmt"



func variable() {
	// var name string ini sudah tidak di rekomendasi kan lagi

// var	name = "Andi"
// 	fmt.Println(name)

// 	name = "Andi Lauw"
// 	fmt.Println(name)	
name := "Andi"
	fmt.Println(name)
name = "Andi Lauw"
	fmt.Println(name)

	// multiple variable
	var (
		firstname = "Andi"
		lastname = "Lauw"
	)

	fmt.Println(firstname + " " + lastname)
	// fmt.Println(lastname)
}