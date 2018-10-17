package searching

import (
	"fmt"
	"math"
)

func binarySearch(data []int, searchNum int) {
	for i := 0; i < len(data); i++ {
		var first = i
		var last = len(data) - i
		var found = false
		for first <= last && !found {
			var midpoint = int(math.Abs(float64(last+first)) / 2)
			if data[midpoint]+data[i] == 8 {
				fmt.Printf("%d, %d", data[i], data[midpoint])
				found = true
			} else {
				if data[i]+data[midpoint] < searchNum {
					first = midpoint + 1
				} else {
					last = midpoint - 1
				}
			}
		}
	}
}
