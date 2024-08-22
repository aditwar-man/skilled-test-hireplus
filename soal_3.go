package main

import (
	"fmt"
)

func isValidString(s string) bool {
	// Panjang string harus antara 1 sampai 4096
	if len(s) < 1 || len(s) > 4096 {
		return false
	}

	// Peta tanda kurung pembuka ke penutup
	opening := map[rune]rune{'<': '>', '{': '}', '[': ']'}
	closing := map[rune]rune{'>': '<', '}': '{', ']': '['}

	// Stack untuk menyimpan tanda kurung pembuka
	stack := []rune{}

	for _, char := range s {
		if _, exists := opening[char]; exists {
			// Jika tanda kurung pembuka, masukkan ke dalam stack
			stack = append(stack, char)
		} else if matchingOpening, exists := closing[char]; exists {
			// Jika tanda kurung penutup, cek apakah stack kosong atau tanda terakhir tidak sesuai
			if len(stack) == 0 || stack[len(stack)-1] != matchingOpening {
				return false
			}
			// Hapus tanda kurung pembuka dari stack
			stack = stack[:len(stack)-1]
		} else {
			// Jika karakter yang tidak valid ditemukan
			return false
		}
	}

	// Memastikan tidak ada tanda kurung pembuka yang belum ditutup
	return len(stack) == 0
}

func main() {
	fmt.Println(isValidString("{{[<>[{{}}]]}}")) // TRUE
	fmt.Println(isValidString("]<[{}}"))         // FALSE
}
