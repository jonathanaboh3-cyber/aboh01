package piscine

func Unmatch(a []int) int {
	count := make(map[int]int)

	for _, num := range a {
		count[num]++
	}

	for _, num := range a {
		if count[num]%2 != 0 {
			return num
		}
	}

	return -1
}
