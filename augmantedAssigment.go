package main

import "fmt"

func augmantedAssignment() {

	// var i = "*"
	// i += strings.Repeat("*", 10)
	// i += "\n"
	// i += strings.Repeat("*",9)
	// i += "\n"
	// i += strings.Repeat("*",8)

	// fmt.Println(i)

	// i += 10
	// fmt.Println(i)

	for i := 1; i <= 5; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
	
		fmt.Println()
	}
fmt.Println("\n")

for i :=5; i >1; i--{
	for j :=0; j<i; j++{
		print("*")
	}
	for j:=0;j<(5-i)*2;j++{
		print(" ")
	}
	for j :=0; j<i; j++{
		print("*")
	}
	fmt.Println()
}

for i :=1; i <=5; i++{
	for j :=0; j<i; j++{
		print("*")
	}
	for j:=0;j<(5-i)*2;j++{
		print(" ")
	}
	for j :=0; j<i; j++{
		print("*")
	}
	fmt.Println()
}
	// for i := 1; i <= 5; i++ {
	// 	for j := 0; j < i; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }
// augmentedAssignment()
var i = 10
i += 5
fmt.Println(i)

var j = 2
j++
fmt.Println(j)
j--
fmt.Println(j)
}