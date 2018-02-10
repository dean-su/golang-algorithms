package _1_test

import (
	"fmt"
	"time"
	"math/rand"
)


func mergeSort(arrayToSort []int) []int {

	var left, right, sortedArray []int
	var leftSize, rightSize int
	var midIndex, currentLeftIndex, currentRightIndex int
	var currentSortdIndex int
	arrayLength := len(arrayToSort)
	if arrayLength < 2 {
		return arrayToSort
	}

	// STEP 1: divide the array in half
	midIndex = arrayLength/2
	if midIndex > 0 {
		leftSize = midIndex
	}else{
		leftSize = 1
	}
	left = arrayToSort[:midIndex]
	//leftSize = midIndex

	right = arrayToSort[midIndex:]
	rightSize = arrayLength - midIndex
	// STEP 2: sort each half
	mergeSort(left)
	mergeSort(right)

	// STEP 3: merge the sorted halves
	sortedArray = make([]int, arrayLength)
	currentLeftIndex = 0
	currentRightIndex = 0
	/*for currentSortdIndex = 0;
	currentSortdIndex < arrayLength;
	currentSortdIndex++ {
		if currentLeftIndex > leftSize-1 && currentRightIndex < rightSize-1 {
			sortedArray[currentSortdIndex] = right[currentRightIndex]
			currentRightIndex++
		} else if currentRightIndex > rightSize-1 && currentLeftIndex < leftSize-1 {
			sortedArray[currentSortdIndex] = left[currentLeftIndex]
			currentLeftIndex++
		} else if left[currentLeftIndex] < right[currentRightIndex]{
			sortedArray[currentSortdIndex] = left[currentLeftIndex]
			currentLeftIndex++
		} else {
			sortedArray[currentSortdIndex] = right[currentRightIndex]
			currentRightIndex++
		}

	}*/
	for currentSortdIndex = 0; currentSortdIndex < arrayLength; currentSortdIndex++ {
		if currentLeftIndex < leftSize && (currentRightIndex >= rightSize || left[currentLeftIndex] < right[currentRightIndex]) {
			sortedArray[currentSortdIndex] = left[currentLeftIndex]
			currentLeftIndex++
		}else {
			sortedArray[currentSortdIndex] = right[currentRightIndex]
			currentRightIndex++
		}
 	}

	return sortedArray
}

func main() {
	iArray := []int{1,2,3,4,5,6,7,8,10,9}
	fmt.Println("Before sorted:", iArray)
	fmt.Println("After sorted:", mergeSort(iArray))

	fmt.Println("Before sorted:", iArray)
	fmt.Println("After sorted:", MergeSort(iArray))
	unsorted := generateSlice(10)
	fmt.Println("Before sorted:", unsorted)
	fmt.Println("After sorted:", MergeSort(unsorted))
}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(99999) - rand.Intn(99999)
	}
	return slice
}

// Runs MergeSort algorithm on a slice single
func MergeSort(slice []int) []int {

	if len(slice) < 2 {
		return slice
	}
	mid := (len(slice)) / 2
	return merge(MergeSort(slice[:mid]), MergeSort(slice[mid:]))
}

// Merges left and right slice into newly created slice
func merge(left, right []int) []int {

	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}