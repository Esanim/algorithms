package searching

import (
	"fmt"
)

func quadraticSearch(data []int, searchNum int) {
	for i := 0; i < len(data); i++ {
		for j := i; j < len(data)-i; j++ {
			if data[i]+data[j] == searchNum {
				fmt.Printf("%d, %d", data[i], data[j])
			}
		}
	}
}
