package big_o_examples

import "fmt"

func PrintAllPairs(list []int) {
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list); j ++ {
			fmt.Printf("%v:%v", list[i], list[j])
		}
	}
}
