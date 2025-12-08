package main

import "fmt"

func operatorBoolean() {
	var nilai_ujian = 95
	var nilai_absensi = 80

	var lulus = nilai_ujian > 80 && nilai_absensi > 80
	var hasil string
	if lulus {
		hasil = "Lulus"
	} else {
		hasil = "Tidak Lulus"
	}
	fmt.Println(hasil)
}