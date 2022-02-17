package main

import (
	"fmt"
)

func main() {
	fmt.Println(possibleTimes([]int{1, 2, 3, 4})) //Answer = 10; "12:34", "12:43", "13:24", "13:42", "14:23", "14:32", "21:34", "21:43", "23:14", "23:41"
	fmt.Println(possibleTimes([]int{2, 2, 1, 9})) // Answer = 4; "12:29", "19:22", "21:29", "22:19"
}

//Find all possible 24 hour time combinations from the array of integers
func possibleTimes(digits []int) int {
	var permutation func([]int, int)
	count := 0
	exists := make(map[string]bool)

	//Implemetation of permutation using Heap's algorithm
	permutation = func(digits []int, n int) {
		if n == 1 {
			if validateTime(digits, exists) {
				count++
				// printTime(digits)
			}
		} else {
			for i := 0; i < n; i++ {
				permutation(digits, n-1)
				//If n is odd, swap first and last else if n is even, swap ith and last element
				if n%2 == 1 {
					temp := digits[0]
					digits[0], digits[n-1] = digits[n-1], temp
				} else {
					temp := digits[i]
					digits[i], digits[n-1] = digits[n-1], temp
				}
			}
		}
	}
	permutation(digits, len(digits))
	return count
}

func validateTime(time []int, exists map[string]bool) bool {
	timeString := fmt.Sprintf("%d%d:%d%d", time[0], time[1], time[2], time[3])
	if time[0]*10+time[1] < 24 && time[2]*10+time[3] < 60 {
		if _, ok := exists[timeString]; !ok {
			exists[timeString] = true
			return true
		}
		return false
	}
	return false
}

func printTime(time []int) {
	fmt.Println(time[0]*10+time[1], ":", time[2]*10+time[3])
}
