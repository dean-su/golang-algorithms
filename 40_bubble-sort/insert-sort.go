package _0_bubble_sort

import (
	"fmt"
)

func main() {
	aa := [5]int{7, 9, 3, 19, 8}
	fmt.Println(aa)
	for i := 1; i < len(aa); i++ {
		tmp := aa[i]
		for j := i - 1; j >= 0 && aa[j] > tmp; j-- {
			aa[j], aa[j+1] = aa[j+1], aa[j]
		}
	}
	fmt.Println(aa)
}
