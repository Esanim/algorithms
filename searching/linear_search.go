package searching

func linear(data []int, searchNum int) {
	var first, last = 0, len(data) - 1
	for first <= last {
		s := data[first] + data[last]
		if s == searchNum {
			first++
			last--
		} else if s < searchNum {
			first++
		} else {
			last--
		}
	}
}
