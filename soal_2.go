package main

import (
	"fmt"
)

func hitungKembalian(totalBelanja, jumlahBayar int) (int, int, map[int]int, bool) {
	pecahanUang := []int{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}
	kembalian := jumlahBayar - totalBelanja

	if kembalian < 0 {
		return 0, 0, nil, false
	}

	kembalianBulat := (kembalian / 100) * 100

	detailPecahan := make(map[int]int)

	sisaKembalian := kembalianBulat
	for _, pecahan := range pecahanUang {
		if sisaKembalian >= pecahan {
			jumlahPecahan := sisaKembalian / pecahan
			detailPecahan[pecahan] = jumlahPecahan
			sisaKembalian -= jumlahPecahan * pecahan
		}
	}

	return kembalian, kembalianBulat, detailPecahan, true
}

func main() {
	totalBelanja1 := 700649
	jumlahBayar1 := 800000
	kembalian1, kembalianBulat1, detailPecahan1, valid1 := hitungKembalian(totalBelanja1, jumlahBayar1)
	if valid1 {
		fmt.Printf("Kembalian yang harus diberikan kasir: Rp %d, dibulatkan menjadi Rp %d\n", kembalian1, kembalianBulat1)
		fmt.Println("Pecahan uang:")
		for pecahan, jumlah := range detailPecahan1 {
			fmt.Printf("%d lembar/koin Rp %d\n", jumlah, pecahan)
		}
	} else {
		fmt.Println("False, kurang bayar")
	}

	totalBelanja2 := 575650
	jumlahBayar2 := 580000
	kembalian2, kembalianBulat2, detailPecahan2, valid2 := hitungKembalian(totalBelanja2, jumlahBayar2)
	if valid2 {
		fmt.Printf("\nKembalian yang harus diberikan kasir: Rp %d, dibulatkan menjadi Rp %d\n", kembalian2, kembalianBulat2)
		fmt.Println("Pecahan uang:")
		for pecahan, jumlah := range detailPecahan2 {
			fmt.Printf("%d lembar/koin Rp %d\n", jumlah, pecahan)
		}
	} else {
		fmt.Println("False, kurang bayar")
	}

	totalBelanja3 := 657650
	jumlahBayar3 := 600000
	_, _, _, valid3 := hitungKembalian(totalBelanja3, jumlahBayar3)
	if !valid3 {
		fmt.Println("\nOutput: False, kurang bayar")
	}
}
