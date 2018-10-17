package sort

func countingSort(theArray []int32, maxValue int) []int32 {
	var numCounts = make([]int, maxValue+1)

	// populate numCounts
	for _, num := range theArray {
		numCounts[num]++
	}

	// populate the final sorted array
	var sortedArray = make([]int32, len(theArray))
	currentSortedIndex := 0

	for num := 0; num < len(numCounts); num++ {
		count := numCounts[num]

		for i := 0; i < count; i++ {
			sortedArray[currentSortedIndex] = int32(num)
			currentSortedIndex++
		}
	}

	return sortedArray
}
