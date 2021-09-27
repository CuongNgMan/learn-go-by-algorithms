package array_operator

func ReserveArray(arr []int) []int {
	l := len(arr)
	mid :=  l / 2
	for i:=0; i <= mid; i++ {
		arr[i], arr[l - i - 1] = arr[l - i -1], arr[i]
	}
	return arr
}
