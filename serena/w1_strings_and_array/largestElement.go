package main

func FindLargestElement()int {
	input := [7]int{1, 5, 3, 9, 2, 8, 6666}
	largest := 0
	for _, value := range input {
		if value > largest {
			largest = value
		}
	}
	return largest
}