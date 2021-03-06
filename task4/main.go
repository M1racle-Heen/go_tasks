package main

import (
	"fmt"
	"sort"
)

func main() {
	first := NewPerson("Said", "Dias")
	second := NewPerson("Abs", "Abs")
	third := NewPerson("Bob", "Dias")
	fourth := NewPerson("John", "Abs")
	fifth := NewPerson("Jenny", "Dias")
	sixth := NewPerson("Anton", "Abs")

	people := []*Person{
		first,
		second,
		third,
		fourth,
		fifth,
		sixth,
	}

	any := []*Anything{{1}, {2}, {1}, {"Hello"}, {"Hello"}, {1}, {2}, {1}}

	fmt.Println("--------------------------------")

	for _, p := range people {
		fmt.Println(p)
	}

	fmt.Println("--------------------------------")

	sort.Sort(PersonSlice(people))

	for _, p := range people {
		fmt.Println(p)
	}

	fmt.Println("--------------------------------")

	fmt.Println(IsPalindrome(AnySlice(any)))

	fmt.Println("--------------------------------")

	fmt.Println(Fold([]int{1, 2, 3, 4, 5, 6}, 5, func(a int, b int) int {
		return a * b
	}))

	fmt.Println("--------------------------------")
}

// Problem 1: Sorting Names
// Sorting in Go is done through interfaces!
// To sort a collection (such as a slice), the type must satisfy sort.Interface,
// which requires 3 methods: Len() int, Less(i, j int) bool, and Swap(i, j int).
// To actually sort a slice, you need to first implement all 3 methods on a
// custom type, and then call sort.Sort on your the PersonSlice type.
// See the Go documentation: https://golang.org/pkg/sort/ for full details.

// Person stores a simple profile. These should be sorted by alphabetical order
// by last name, followed by the first name, followed by the ID. You can assume
// the ID will be unique, but the names need not be unique.
// Sorting should be case-sensitive and UTF-8 aware.
type Person struct {
	ID        int
	FirstName string
	LastName  string
}

type PersonSlice []*Person

type PersonInterface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// NewPerson is a constructor for Person. ID should be assigned automatically in
// sequential order, starting at 1 for the first Person created.
var id int = 0

func NewPerson(first, last string) *Person {
	id++
	person := &Person{ID: id, FirstName: first, LastName: last}
	// TODO
	return person
}

func (p PersonSlice) Len() int {
	return len(p)
}

func (p PersonSlice) Less(i, j int) bool {
	return p[i].FirstName < p[j].FirstName || p[i].LastName < p[j].LastName
}

func (p PersonSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Problem 2: IsPalindrome Redux
// Using a function that simply requires sort.Interface, you should be able to
// check if a sequence is a palindrome. You may use, adapt, or modify your code
// from HW0. Note that the input does not need to be a string: any type which
// satisfies sort.Interface can (and will) be used to test. This means that the
// only functionality you should use should come from the sort.Interface methods
// Ex: [1, 2, 1] => true
type Anything struct {
	anything interface{}
}
type AnySlice []*Anything
type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func (a AnySlice) Len() int {
	return len(a)
}

func (a AnySlice) Less(i, j int) bool {
	switch a[i].anything.(type) {
	case int:
		return a[i].anything.(int) < a[j].anything.(int)
	case float64:
		return a[i].anything.(float64) < a[j].anything.(float64)
	case string:
		return a[i].anything.(string) < a[j].anything.(string)
	default:
		fmt.Printf("I don't know, ask stackoverflow.")
	}
	return false
}

func (a AnySlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// IsPalindrome checks if the string is a palindrome.
// A palindrome is a string that reads the same backward as forward.
func IsPalindrome(s sort.Interface) bool {
	i, j := 0, s.Len()-1
	for j > i {
		if !s.Less(j, i) && !s.Less(i, j) {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

// Problem 3: Functional Programming
// Write a function Fold which applies a function repeatedly on a slice,
// producing a single value via repeated application of an input function.
// The behavior of Fold should be as follows:
//   - When s is empty, return v (default value)
//   - When s has 1 value (x0), apply f once: f(v, x0)
//   - When s has 2 values (x0, x1), apply f twice, from left to right: f(f(v, x0), x1)
//   - Continue this pattern recursively to obtain the final result.

// Fold applies a left to right application of f on s starting with v.
// Note the argument signature of f - func(int, int) int.
// This means f is a function which has 2 int arguments and returns an int.
func Fold(s []int, v int, f func(int, int) int) int {
	z := 0

	for i := range s {
		if i == 0 {
			z = f(v, s[i])
		} else {
			z = f(z, s[i])
		}
	}
	// TODO
	return z
}
