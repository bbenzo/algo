package fast_slow_pointers

import (
	"log"
	"strconv"
)

func HappyNumber(num string, squareSums []int) bool {
	sum := 0
	for _, val := range digits(num) {
		sum += val * val
	}

	if sum == 1 {
		return true
	}

	squareSums = append(squareSums, sum)

	if hasCycle(squareSums) {
		return false
	}

	return HappyNumber(strconv.Itoa(sum), squareSums)
}

func digits(num string) []int {
	var digits []int

	for i := range num {
		converted, err := strconv.Atoi(string(num[i]))
		if err != nil {
			log.Fatalf("no number %v", converted)
		}

		digits = append(digits, converted)
	}

	return digits
}

func hasCycle(arr []int) bool {
	fast, slow := 0, 0

	for slow < len(arr) {

		if fast == len(arr) - 2 {
			fast = 0
		} else if fast == len(arr) - 1 {
			fast = 1
		} else {
			fast += 2
		}

		slow++

		if slow == fast {
			return true
		}
	}

	return false
}