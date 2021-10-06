package main

import (
	"algo-playground/array_operator"
	"algo-playground/list"
	"algo-playground/searcher"
	"algo-playground/sorter"
	"fmt"
	"math/rand"
	"time"
)

func simpleFun() {
	scores := []int{9, 10, 6, 7, 5, 8, 4}
	length := len(scores)
	for i := 0; i < length; i++ {
		fmt.Printf("%d,", scores[i])
	}
}

func writeToConsole(arr []int) {
	for i := 0; i < len(arr) - 1; i++ {
		fmt.Printf("%d,", arr[i])
	}
	fmt.Println("\n")
}

func main() {
	totalInput := 50
	maxValue :=  50
	inputArr := []int{}

	for i := 0; i <= totalInput; i++ {
		rand.Seed(time.Now().UnixNano())
		inputArr = append(inputArr, rand.Intn(maxValue))
	}
	rand.Shuffle(len(inputArr), func(i, j int) {
		inputArr[i], inputArr[j] = inputArr[j], inputArr[i]
	})

	result := array_operator.FindMax(inputArr)
	fmt.Println("Max value: ", result)
	result = array_operator.FindMin(inputArr)
	fmt.Println("Min value: ", result)

	newArr := make([]int, len(inputArr) + 1)
	array_operator.LinearArrayInsert(inputArr, newArr, 10, -1)
	fmt.Println("New inserted array")
	writeToConsole(newArr)

	newArr = make([]int, len(inputArr) - 1)
	indexToBeDeleted := 1
	array_operator.LinearArrayDelete(inputArr, newArr, indexToBeDeleted)
	fmt.Printf("New array after index %d is deleted\n", 1)
	writeToConsole(newArr)

	fmt.Println("New reversed array")
	array_operator.ReserveArray(newArr)
	writeToConsole(newArr)

	/*
	* Sorting Algorithms
	*/
	var sortedArray []int
	fmt.Println("Unsort Array:")
	writeToConsole(inputArr)

	//sortedArray = sorter.BubbleSort(inputArr)
	//sortedArray = sorter.SelectSort(inputArr)
	//sortedArray = sorter.InsertSort(inputArr)
	sortedArray = sorter.ShellSort(inputArr)
	writeToConsole(sortedArray)

	/*
	* Searching Algorithms
	*/
	var shouldAsk = false
	var value int = 0
	fmt.Println("Enter value to search:\n")
	if shouldAsk {
		fmt.Scan(&value)
	}

	//foundIndex := searcher.LinearSearch(inputArr, value)
	foundIndex := searcher.BinarySearch(inputArr, value)
	if foundIndex != -1 {
		fmt.Printf("Found %d at index %d", value, foundIndex)
	} else {
		fmt.Printf("Value %d not found", value)
	}
	fmt.Println()

	/*
	* Data structure - Single LinkedList
	*/
	fmt.Println("Linked-List")
	ll := list.NewLinkedList()
	ll.Init([]string{"TpHCM", "Ha Noi"})
	ll.Output()
	ll.AppendByValue("Sapa")
	ll.Output()
	ll.Insert(1,"Nha Trang")
	ll.Output()
	fmt.Println("After delete")
	ll.Delete(0)
	ll.Output()

	/*
	* Data structure - Doubly LinkedList
	*/
	fmt.Println("Doubly Linked-List")
	dl := list.NewDoublyLinkedList()
	dl.Init([]string{"America", "Berlin", "Congo", "Denmark", "England", "France"})
	dl.Insert("Vietnam", 2)
	dl.Insert("Russia", 8)
	dl.Delete(0)
	dl.OutputFromHead()
	dl.OutputFromTail()

	/*
	* Data structure - One-way Circular LinkedList
	*/
	fmt.Println("One-way Circular LinkedList")
	cLL := list.NewOneWayCirularLinkedList()
	cLL.Init([]string{"A", "B", "C", "D"})
	cLL.Insert(1, "G")
	cLL.Insert(4, "E")
	cLL.Insert(5, "F")

	_, err := cLL.Delete(3)
	if err != nil {
		fmt.Println(err)
		return
	}
	cLL.OutputFromHead()
	cLL.OutputFromTail()
}
