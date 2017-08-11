// data structures
// slice = list ~= array
// map = dictionary = key:value
// struct = different pieces of data = has properties and methods
// struct properties = fields

package main

import "fmt"

func main() {
	array()
	slice()
	mapExample()
	differentMaps()
}

func array() {
	/*
		numbers in brackets = array
		no numbers = slice
	*/
	var array [5]int // number - how many items
	var slice []int
	fmt.Println(len(array))
	fmt.Println(len(slice))

	for i := 0; i < len(array); i++ {
		array[i] = i
	}
	fmt.Println(array)

}

func slice() {
	/*
		make(<type>,<length>,<capacity>)
	*/
	slice := make([]int, 5, 10)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}

func mapExample() {
	m := make(map[string]int)
	m["key1"] = 42
	fmt.Println(m)
}

func differentMaps() {
	var map1 = make(map[int]int)
	fmt.Println(map1 == nil)
	map1[1] = 1
	map1[2] = 2
	fmt.Println(map1)

	var map2 map[int]int // kinda useless, no append
	// map2[1] = 1    // error
	fmt.Println(map2)
	fmt.Println(map2 == nil)

	map3 := map[int]int{1: 1, 2: 2}
	delete(map3, 1)
	fmt.Println(map3)

	if val, exists := map3[2]; exists {
		// delete(map3, 3)
		fmt.Println(val)
		fmt.Println(exists)
	} else {
		fmt.Println("Doesn't exist")
	}
}
