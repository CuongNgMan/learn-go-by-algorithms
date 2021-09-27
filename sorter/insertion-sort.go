package sorter

import "fmt"

func InsertSort(arr []int) []int {
	fmt.Println("Insertion Sort")
	l := len(arr)
	for i:=0;i < l;i++ {
		insertElement := arr[i]
		insertPosition := i
		for j := insertPosition - 1; j >= 0; j-- {
			if arr[j] > insertElement {
				arr[j+1] = arr[j]
				insertPosition--
			}
		}
		arr[insertPosition] = insertElement
	}
	return arr
}
