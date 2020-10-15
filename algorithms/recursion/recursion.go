package recursion

func IterativeFactorial(num int) int {
	result := 1

	for num > 0 {
		result *= num
		num--
	}

	return result
}

func RecursiveFactorial(num int) int {
	if num < 2 {
		return num
	}
	return num * RecursiveFactorial(num - 1)
}

func IsPalindrome(word string) bool {
	if len(word) < 2 {
		return true
	}

	if word[0] == word[len(word)-1] {
		return IsPalindrome(word[1:len(word)-1])
	}

	return false
}

func RecursivePow(x, n int) int {
	if n == 1 {
		return x
	}

	if n < 0 || n % 2 != 0 {
		return -1
	}

	return RecursivePow(x, n/2) * RecursivePow(x, n/2)
}
