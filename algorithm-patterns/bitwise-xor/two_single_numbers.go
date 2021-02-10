package bitwise_xor

func TwoSingleNumbers(arr []int) []int {
	n1xn2 := 0
	for i := range arr {
		n1xn2 ^= arr[i]
	}

	rightMostSetBit := 1
	for (rightMostSetBit & n1xn2) == 0 {
		rightMostSetBit <<= 1
	}

	num1, num2 := 0, 0
	for i := range arr {
		if (arr[i] & rightMostSetBit) != 0 {
			num1 ^= arr[i]
		} else {
			num2 ^= arr[i]
		}
	}

	return []int{num1, num2}
}
