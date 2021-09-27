package searcher

func LinearSearch(arr []int, searchValue int) int  {
	l := len(arr)
	for i:= 0; i< l;i++ {
		if arr[i] == searchValue {
			return i
		}
	}
	return -1
}