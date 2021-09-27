package array_operator

func LinearArrayDelete(origArray []int, newArray []int, index int) error {
	origLen := len(origArray)
	for i := 0; i < origLen; i++ {
		if i < index {
			newArray[i] = origArray[i]
		}
		if i > index {
			newArray[i - 1] = origArray[i]
		}
	}
	return nil
}
