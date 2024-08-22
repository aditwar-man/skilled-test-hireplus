package main

import (
	"fmt"
	"strings"
)

func findMatchingStrings(N int, strs []string) interface{} {
	isEqual := func(str1, str2 string) bool {
		if len(str1) != len(str2) {
			return false
		}
		for i := 0; i < len(str1); i++ {
			if strings.ToLower(string(str1[i])) != strings.ToLower(string(str2[i])) {
				return false
			}
		}
		return true
	}

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if isEqual(strs[i], strs[j]) {
				matchingIndices := []int{i + 1, j + 1}
				for k := j + 1; k < N; k++ {
					if isEqual(strs[i], strs[k]) {
						matchingIndices = append(matchingIndices, k+1)
					}
				}
				return matchingIndices
			}
		}
	}

	return false
}

func main() {
	inputs := [][]string{
		{"abcd", "acbd", "aaab", "acbd"},
		{"Satu", "Sate", "Tujuh", "Tusuk", "Tujuh", "Sate", "Bonus", "Tiga", "Puluh", "Tujuh", "Tusuk"},
		{"pisang", "goreng", "enak", "sekali", "rasanya"},
	}

	for _, strs := range inputs {
		N := len(strs)
		result := findMatchingStrings(N, strs)
		switch v := result.(type) {
		case []int:
			for _, idx := range v {
				fmt.Print(idx, " ")
			}
			fmt.Println()
		case bool:
			fmt.Println("false")
		}
	}
}
