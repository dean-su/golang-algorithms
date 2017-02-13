package main

import (
	"fmt"
)

func bubble_sort(yarray []int) {

	n := len(yarray)

	for j := 0; j < n-1; j++ {
		for i := 0; i < n-1-j; i++ {
			if yarray[i] > yarray[i+1] {
				yarray[i], yarray[i+1] = yarray[i+1], yarray[i]
			}

		}
	}
}

func main() {

	xarray := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	fmt.Println("before sort:", xarray)
	bubble_sort(xarray)
	fmt.Println("after sort:", xarray)
}
