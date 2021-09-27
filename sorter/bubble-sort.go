package sorter

import "fmt"

func BubbleSort(arr []int) []int {
	fmt.Println("Bubble Sort")
	l := len(arr)
	for i:=0; i < l - 1; i++ {
		for  j:=i+1; j < l - 1; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
