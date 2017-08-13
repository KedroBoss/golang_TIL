package main

import "fmt"
import "sort"

type people []string

// ByName type is allies to people, which is an allies to slice of strings
// It implements Len, Swap, Less methods -> implements Interface interface
type ByName people

func (a ByName) Len() int      { return len(a) }
func (a ByName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// The length must be at least 2 characters long
func (a ByName) Less(i, j int) bool { return a[i][:2] < a[j][:2] }

func main() {
	group := people{"Bb", "Ba", "Ab", "Di", "Bc"}
	s := []string{"Ba", "Ab", "Di", "Bc"}
	i := []int{1, 5, 8, 2, 9, 0, 22}

	// Strings takes a slice of strings
	sort.Strings(s)
	fmt.Println(s)

	// Sort takes Interface implementation
	// ByName implements Interface
	// ByName is allies to people
	// Group is a slice of people
	sort.Sort(ByName(group))
	fmt.Println(group)

	// Sort takes Interface
	// Reverse takes Interface and returns Interface
	// IntSlice takes a slice of integers
	// IntSlice implements Interface
	sort.Sort(sort.Reverse(sort.IntSlice(i)))
	fmt.Println(i)
}
