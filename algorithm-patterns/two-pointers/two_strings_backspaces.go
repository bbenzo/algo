package two_pointers

func TwoStringsBackspaces(str1, str2 string) bool {

	ptr1 := len(str1) - 1
	ptr2 := len(str2) - 1
	for ptr1 > 0 {
		if string(str1[ptr1]) == "#" {
			backspaces := 0
			for string(str1[ptr1]) == "#" {
				backspaces++
				ptr1--
			}

			ptr1 -= backspaces
		}

		if string(str2[ptr2]) == "#" {
			backspaces := 0
			for string(str2[ptr2]) == "#" {
				backspaces++
				ptr2--
			}

			ptr2 -= backspaces
		}

		if str1[ptr1] != str2[ptr2] {
			return false
		}

		ptr1--
		ptr2--
	}

	return true
}
