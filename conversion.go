package main

import "fmt"

func main() {
	// var nilai32 int32 = 32768
	// var nilai64 int64 = int64(nilai32)
	// var nilai16 int16 = int16(nilai32)

	// fmt.Println(nilai32)
	// fmt.Println(nilai64)
	// fmt.Println(nilai16)

	fmt.Println("Sebelum Di konversi")
		fmt.Println(len("Andi Lauw"))
		fmt.Println("Andi Lauw"[0])
			fmt.Println("Andi Lauw"[1])


	var name = "Andi Lauw"
	// var name2 uint8 = name[0]
	// var name2String = string(name2)
	// var name3 uint8 = name[1]
	// var name3String = string(name3)
	// 	fmt.Println(("Setelah DI konversi"))	
	// fmt.Println(len(name))
	// 	fmt.Println(name2String)
	// 		fmt.Println(name3String)

	name2 := name[0] 
	name3 := name[1] 

	name4 := string(name2)
	name5 := string(name3)

	fmt.Println(("Setelah DI konversi"))	
	fmt.Println(len(name))
		fmt.Println(name4)
			fmt.Println(name5)
}