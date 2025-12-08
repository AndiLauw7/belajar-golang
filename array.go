package main

import "fmt"

func main() {
	var name [3]string

	name[0] = "Andi"
	name[1] = "Lauw"
	name[2] = "Lauw"
	fmt.Println(name[0])
	
	var nilaiArray = [3] string{
"andi",
"safa",
"riansyah",

	}
	fmt.Println(nilaiArray)

	var panjangArray=[...]string{
"",
"",
"",
"",
"",
	}
	fmt.Println("nilai array",panjangArray, "panjang aray",len(panjangArray))
}