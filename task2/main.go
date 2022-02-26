package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(ParsePhone("1 2 3 4 5 6 7 8 9 0"))

	fmt.Println(Anagram("йцукенгшщзхъфывапролджэячсмитьбю", "юэъхжбдзьлщтоширгмпнсаечвкяыуфцй"))

	a := []int{1, 3, 4, 5, 6, 7, 8, 2, 2, 2, 2, 6}
	fmt.Println(FindEvens(a))

	fmt.Println(SliceProduct(a))

	fmt.Println(Unique(a))

	b := make(map[string]int)
	b["b1"] = 10
	b["b2"] = 20
	fmt.Println(InvertMap(b))

	fmt.Println(TopCharacters("банан", 2))
}

// ParsePhone parses a string of numbers into the format (123) 456-7890.
// This function should handle any number of extraneous spaces and dashes.
// All inputs will have 10 numbers and maybe extra spaces and dashes.
// For example, ParsePhone("123-456-7890") => "(123) 456-7890"
//              ParsePhone("1 2 3 4 5 6 7 8 9 0") => "(123) 456-7890"
func ParsePhone(phone string) string {
	// TODO

	str := "("
	s := strings.Replace(phone, " ", "", -1)
	s = strings.Replace(s, "-", "", -1)
	for i := 0; i < len(s); i++ {
		if i == 3 {
			str += ") "
		}
		if i == 6 {
			str += " "
		}
		str = str + string(s[i])
	}

	return str
}

// Anagram tests whether the two strings are anagrams of each other.
// This function is NOT case sensitive and should handle UTF-8
func Anagram(s1, s2 string) bool {
	// TODO

	if len(s1) != len(s2) {
		return false
	}
	n1 := strings.Split(s1, "")
	sort.Strings(n1)
	n2 := strings.Split(s2, "")
	sort.Strings(n2)

	for i := 0; i < len(n1); i++ {
		if n1[i] != n2[i] {
			return false
		}
	}

	return true

}

// FindEvens filters out all odd numbers from input slice.
// Result should retain the same ordering as the input.
func FindEvens(e []int) []int {
	// TODO

	result := make([]int, 0)
	for _, item := range e {
		if item%2 == 0 {
			result = append(result, item)
		}
	}
	return result
}

// SliceProduct returns the product of all elements in the slice.
// For example, SliceProduct([]int{1, 2, 3}) => 6
func SliceProduct(e []int) int {
	// TODO

	result := 0
	for _, item := range e {
		result += item
	}

	return result
}

// Unique finds all distinct elements in the input array.
// Result should retain the same ordering as the input.
func Unique(e []int) []int {
	// TODO
	result := make([]int, 0)

	for _, item := range e {
		b := true
		for _, item2 := range result {
			if item2 == item {
				b = false
			}
		}
		if b {
			result = append(result, item)
		}
	}

	return result
}

// InvertMap inverts a mapping of strings to ints into a mapping of ints to strings.
// Each value should become a key, and the original key will become the corresponding value.
// For this function, you can assume each value is unique.
func InvertMap(kv map[string]int) map[int]string {
	// TODO
	result := make(map[int]string)

	for k, v := range kv {
		result[v] = k
	}

	return result
}

// TopCharacters finds characters that appear more than k times in the string.
// The result is the set of characters along with their occurrences.
// This function MUST handle UTF-8 characters.
func TopCharacters(s string, k int) map[rune]int {
	// TODO
	result := make(map[rune]int)

	for _, item := range s {
		result[item] += 1
	}
	for key, value := range result {
		if value < k {
			delete(result, key)
		}
	}
	return result
}
