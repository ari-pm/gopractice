package main

import (
	"time"
	"fmt"
	"math/rand"
)

func BubbleSort(arr []int) {
	// We compare each 2 elements next to each otehr, starting from the beginning, and swap if the left is larger than the rigth.
	// At each round the largest element will be on the right most side.
	// In the next round we only find the largest among the rest of the elements.
	for i:=len(arr)-1; i>1; i-- {
		for j:=0; j<i; j++ {
			if arr[j]>arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func InsertionSort(arr []int) {
	// Sort elements until i, starting from i=2
	// Move by 1 at each step towards right
	// At each step compare the new element on the right with the elements on the left (if smaller, then move the element on the left one to the right) until finding the first element smaller.
	// Then you can replace that last index with the new element.
	for i:=1; i<len(arr); i++ {
		val := arr[i]
		var j int
		for j=i-1; j>=0 && arr[j] > val; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = val
	}
}
/* Function to sort an array using insertion sort
void insertionSort(int arr[], int n)
{
   int i, key, j;
   for (i = 1; i < n; i++)
   {
       key = arr[i];
       j = i-1;
 
        // Move elements of arr[0..i-1], that are
    	// greater than key, to one position ahead
        //  of their current position 
       while (j >= 0 && arr[j] > key)
       {
           arr[j+1] = arr[j];
           j = j-1;
       }
       arr[j+1] = key;
   }
}
*/

// Here the assumption is that the array will be sorted, so no other output.
func QuickSort(arr []int) {
	//Pivot is the last element
	//Push everything smaller/equal than pivot to the left side
	//At the end swap the last element with the pivot.
	//Pass the left and right side to the QuickSort().
	if len(arr) <= 1 {
		return
	}
	var pivot int
	pivot = arr[len(arr)-1]
	//fmt.Println("pivot", pivot)
	j := 0
	for i:=0; i<len(arr)-1; i++ {
		if arr[i] <= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			j++
		}
	}
	arr[len(arr)-1] = arr[j]
	arr[j] = pivot
	QuickSort(arr[:j])
	QuickSort(arr[j+1:])
}
/*
func MergeSort(arr []int) {
	//Partition
	//Apply mergeSort to each part
	//Merge them
	if len(arr) <= 1 {
		return
	} 
	midIdx := len(arr) / 2
	MergeSort(arr[:midIdx])
	MergeSort(arr[midIdx:])

	arr = mergeMS(arr[:midIdx], arr[midIdx:])	// even if the mergeMS doesn't do the merging inplace, we can fake it.
}
func mergeMS(arr1 []int, arr2 []int) {
	// We know that the first array is the first portion of the total array and the 2nd is the 2nd portion.
	arr1Cpy := make([]int, len(arr1))
	copy(arr1Cpy, arr1)

}
*/

// The following works but I think it is not efficient.  (what is time and memory complexity?)
func MergeSort(arr []int) []int {
	//Partition
	//Apply mergeSort to each part
	//Merge them
	if len(arr) <= 1 {
		return arr
	} 
	midIdx := len(arr) / 2
	arr1 := MergeSort(arr[:midIdx])
	arr2 := MergeSort(arr[midIdx:])

	arr3 := mergeMS(arr1, arr2)	// even if the mergeMS doesn't do the merging inplace, we can fake it.
	return arr3
}
func mergeMS(arr1 []int, arr2 []int) []int {
	// We know that the first array is the first portion of the total array and the 2nd is the 2nd portion.
	arr := make([]int, len(arr1)+len(arr2))
	i, j, idx := 0, 0, 0
	for i<len(arr1) && j<len(arr2){
		if arr1[i] <= arr2[j] {
			arr[idx] = arr1[i]
			i++
		} else {
			arr[idx] = arr2[j]
			j++
		}
		idx++
	}
	for i<len(arr1) {
		arr[idx] = arr1[i]
		idx++
		i++
	} 
	for j<len(arr2) {
		arr[idx] = arr2[j]
		idx++
		j++
	} 

	return arr
}

func HeapSort(arr []int) {
	// I think it works by keeping data in a heap.  Then for sorting, it pulls data from the root (smallest), and reorder
	return
}

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	arr := make([]int, 20)
	for i := range arr {
		arr[i] = rand.Intn(30)
	}

	fmt.Println()
	fmt.Println(arr)
	fmt.Println("BubbleSort")
	BubbleSort(arr)
	fmt.Println(arr)

	for i := range arr {
		arr[i] = rand.Intn(30)
	}
	fmt.Println()
	fmt.Println(arr)
	fmt.Println("InsertionSort")
	InsertionSort(arr)
	fmt.Println(arr)

	for i := range arr {
		arr[i] = rand.Intn(30)
	}
	fmt.Println()
	fmt.Println(arr)
	fmt.Println("QuickSort")
	QuickSort(arr)
	fmt.Println(arr)

	for i := range arr {
		arr[i] = rand.Intn(30)
	}
	fmt.Println()
	fmt.Println(arr)
	fmt.Println("MergeSort")
	//MergeSort(arr)
	fmt.Println(MergeSort(arr))

	for i := range arr {
		arr[i] = rand.Intn(30)
	}
	fmt.Println()
	fmt.Println(arr)
	fmt.Println("HeapSort")
	HeapSort(arr)
	fmt.Println(arr)
}
