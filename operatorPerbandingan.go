package main

import "fmt"

func operatorPerbandingan() {
	const name = "Andi"
	const name2 = "Andi"

	var result = name == name2
	fmt.Println(result)

	const lima = 5
	const enam = 6
	var result2 = lima < enam
	fmt.Println(result2)

	const tiga = 4
	const tiga2 = 3
	var result3 = tiga > tiga2
	fmt.Println(result3)

	const empat = 4
	const empat2 = 4
	var result4 = empat != empat2
	fmt.Println(result4)

}