package big_o

import "fmt"

func PrintEveryItem(list []int) {
	for i := 0; i < len(list); i++ {
		fmt.Println(list[i])
	}
}
