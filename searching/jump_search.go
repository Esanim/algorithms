package searching

import "fmt"
import "math"

func jumpSearch(arr []int, key int) int {
	//block size to jump
	sz := len(arr)
	fmt.Printf("len = %v\n", sz)
	step := int(math.Sqrt(float64(sz)))
	fmt.Printf("step = %v\n", step)
	prev := 0
	//finding the block
	for arr[int(math.Min(float64(step), float64(sz)))-1] < key {
		prev = step
		step += int(math.Sqrt(float64(sz)))
		fmt.Printf("new step = %v\n", step)
		if prev >= sz {
			return -1
		}
	}
	//linear search the block
	for arr[prev] < key {
		prev++
		if prev == int(math.Min(float64(step), float64(sz))) {
			return -1
		}
	}
	if arr[prev] == key {
		return prev
	}
	return -1
}
