package main

func SeparateElements() []int{
	input := []int{1,2,3,4,6,8,7,12}
	var left int = 0
	var right int = len(input) - 1
	for left < right {
		for left < right  && input[left]%2 == 0 {
			left ++
		}
		for left < right && input[right]%2 == 1 {
			right --
		}
		if left < right {
			tmp := input[left]
			input[left] = input[right]
			input[right] = tmp
			left++
			right--
		}
	}
	return input
}