package array_operator

func FindMin(arr []int) int {
	l := len(arr)
	var result int = arr[0]
	for i := 0; i < l-1; i++ {
		if arr[i] < result {
			result = arr[i]
		}
	}
	return result
}
