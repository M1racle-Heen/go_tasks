package main

import "fmt"

func main() {
	fbInt := 15
	fmt.Println(FizzBuzz(fbInt))
	isPrimInt := 5
	fmt.Println(IsPrime(isPrimInt))
	isPalString := "shaahs"
	fmt.Println(IsPalindrome(isPalString))
}

func FizzBuzz(n int) string {
	k := ""
	if n%3 == 0 {
		k = "Fizz"
	} else if n%5 == 0 {
		k = "Buzz"
	}
	if n%5 == 0 && n%3 == 0 {
		k = "FizzBuzz"
	}
	return k
}

func IsPrime(n int) bool {
	if n > 1 {
		for i := 2; i < n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}
	return false
}

func IsPalindrome(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
