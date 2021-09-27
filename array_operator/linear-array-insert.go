package array_operator

import (
	"errors"
)

func LinearArrayInsert(arr []int, tempArr []int, insertIndex int, value int) error {
	origL := len(arr)
	newArrL := len(tempArr)
	if insertIndex < 0 || insertIndex > newArrL-1 {
		return errors.New("Invalid index")
	}
	for i := 0; i < origL; i++ {
		if i < insertIndex {
			tempArr[i] = arr[i]
		} else {
			tempArr[i+1] = arr[i]
		}
	}
	tempArr[insertIndex] = value

	return nil
}
