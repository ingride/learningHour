package main

import "fmt"

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

func maxArray(nums []int) (int){
	var max = MinInt
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max
}
func main() {
	nums := []int{-2, -3, -4, -2}
	fmt.Println("max in array: ", maxArray(nums))
}