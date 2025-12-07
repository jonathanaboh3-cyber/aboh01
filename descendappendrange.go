package piscine

func DescendAppendRange(max, min int) []int {
	if max <= min {
		return []int{} // return an empty, non-nil slice
	}

	result := []int{} // initialize as an empty slice
	for i := max; i > min; i-- {
		result = append(result, i)
	}
	return result
}
