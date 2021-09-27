package searcher

// Array must be sorted
func BinarySearch(arr []int, searchValue int) int {
	l := len(arr)
	low := 0
	mid := 0
	high := l
	for {
		if low >= high {
			break;
		}
		mid = (low + high) / 2
		if arr[mid] == searchValue {
			return mid
		} else if searchValue < arr[mid] {
			high = mid - 1
		} else if arr[mid] < searchValue {
			low = mid + 1
		}
	}

	return -1
}
