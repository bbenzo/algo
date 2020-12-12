package two_pointers

import "sort"


func UniqueTripletsThatAddUpToZero(input []int) [][]int {
	var result [][]int
	start, end := 0, len(input) - 1
	sort.Ints(input) // sorted = -3, -2, -1, 0, 1, 1, 2

	for start < end {
		sum := input[start] + input[end]
		if sum < 0 {
			i := end
			for i > start {
				if sum + input[i] == 0  {
					result = append(result, []int{input[start], input[end], input[i]})
					break
				}
				i--
			}

			for input[start+1] == input[start] {
				start++
			}
			start++
		} else {
			i := start
			for i < end {
				if sum + input[i] == 0  {
					result = append(result, []int{input[start], input[end], input[i]})
					break
				}
				i++
			}

			for input[end-1] == input[end] {
				end--
			}
			end--
		}
	}

	return result
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
