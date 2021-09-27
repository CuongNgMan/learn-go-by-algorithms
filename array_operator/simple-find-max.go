package array_operator

func FindMax(numbs []int) int {
	l := len(numbs)
	copyNumbs := make([]int, l)
	copy(copyNumbs, numbs)
	for i := 0; i < l-1; i++ {
		if copyNumbs[i] > copyNumbs[i+1] {
			copyNumbs[i], copyNumbs[i+1] = copyNumbs[i+1], copyNumbs[i]
		}
	}
	maxValue := copyNumbs[l-1]
	return maxValue
}