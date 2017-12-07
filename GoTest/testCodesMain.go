package main

import "fmt"

var data int
var datap *float64
var arr [10]int
var slic []int

func main() {

	slic = make([]int, 5)
	slic1 := make([]int, 20) // using make to declare and initiate a slice
	for i := 0; i < 20; i++ {
		slic1[i] = i + 1
	}
	copy(slic, slic1[:10]) // using copy for slice

	map1 := make(map[string]bool) // MAP definition

	map1["arezou"] = true
	map1["ryan"] = true
	map1["bird"] = false

	for m := range map1 { // pay attention to the syntax of range here
		fmt.Println(m, map1[m])
	}

	fmt.Println("arr", arr)
	fmt.Println("slic", slic)
	fmt.Println("slic1", slic1)
	fmt.Println("map1", map1)
}
