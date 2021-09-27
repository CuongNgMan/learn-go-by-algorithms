package sorter

import "fmt"

func FancySwap(arr []int, a int, b int) {
	arr[a] = arr[a] + arr[b]
	arr[b] = arr[a] - arr[b]
	arr[a] = arr[a] - arr[b]
}

// This for-loop is required to make the entire array (by current gap) is sorted
func ShellSort(arr []int) []int {
	fmt.Println("Shell Sort")
	l := len(arr)
	for gap:=l/2; gap > 0; gap = gap / 2 {
		for i := gap ; i < l; i++ {
			var j = i
			for {
				if j - gap < 0 || arr[j] >= arr[j-gap] {
					break;
				}
				arr[j], arr[j-gap] = arr[j-gap], arr[j]
				j = j - gap
			}
		}
	}
	return arr
}
