package main

import "fmt"

func typeDeclaration() {
	type NoKTP string

	var noKtpAndi NoKTP = "123456789"

	var ktpContoh string ="11111111111111"
	var hasilContoh NoKTP = NoKTP(ktpContoh) 
	fmt.Println(noKtpAndi)
	fmt.Println(hasilContoh)
}