package main

import "fmt"

func oddAndEven(nums []int) (odd []int, even []int) {
	for _, n := range nums {
		if (n % 2 == 0) {
			even = append(even, n)
		} else {
			odd = append(odd, n)
		}
	}

	return odd, even
}

func main(){

	nums := []int{5, 3, 4, 2}

	odd, even := oddAndEven(nums)

	result := append(even, odd...)

	for _, n := range result {
		fmt.Print(" ", n)
	}


}