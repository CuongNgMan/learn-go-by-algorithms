package sorter

import "fmt"

func SelectSort(arr []int) []int {
	fmt.Println("Selection Sort")
	l := len(arr)
	numbs := make([]int, l)
	copy(numbs, arr)
	var minIndex int
	for i := 0; i < l-1; i++ {
		minIndex = i
		for j := i+1; j<l-1;j++{
			if numbs[j] < numbs[minIndex] {
				minIndex = j
			}
		}
		if i != minIndex {
			numbs[i], numbs[minIndex] = numbs[minIndex], numbs[i]
		}
	}
	return numbs
}
