package main

import (
	"fmt"
)
// A recursive binary search function. It returns location of x in
// given array arr[l..r] is present, otherwise 0
func BinarySearchRecursive(iArray []int, l int, r int, target int) int {
	for r >= l {
		mid := l + (r-l)/2

		if iArray[mid] == target {
			return mid
		}

		if iArray[mid] > target {
			return BinarySearchRecursive(iArray, l, mid, target)
		}
		return BinarySearchRecursive(iArray, mid, r, target)
	}
	return 0
}

func BinarySearchIterative(iArray []int, l int, r int, target int) int  {
	for r >= l {

		mid := l + (r-l)/2

		if iArray[mid] == target {
			return mid
		}

		if iArray[mid] > target {
			r = mid-1
		}else {
			l = mid+1
		}
	}

	return 0
}

func main() {
	iSortedArray := []int{1,4,6,8,13,16}
	iLength := len(iSortedArray)
	iTarget := 8
	fmt.Println("Recursive target:", iTarget," index:", BinarySearchRecursive(iSortedArray, 0, iLength, iTarget), "in ",iSortedArray)
	fmt.Println("Iterative target:", iTarget," index:", BinarySearchIterative(iSortedArray, 0, iLength, iTarget), "in ",iSortedArray)
}
