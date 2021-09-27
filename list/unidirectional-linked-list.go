package list

import (
	"errors"
	"fmt"
)

type LinkedList struct {
	Head   *Node
	Tail   *Node
	length int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head:   nil,
		Tail:   nil,
		length: 0,
	}
}

func (ll *LinkedList) Init(destionations []string) {
	for _, value := range destionations {
		if ll.Head == nil && ll.Tail == nil {
			firstNode := &Node{
				data: value,
				next: nil,
			}
			ll.Head = firstNode
			ll.Tail = firstNode
			ll.length = 1
		} else {
			ll.AppendByValue(value)
		}
	}
}

func (ll LinkedList) Output() {
	var p = ll.Head
	for {
		if p == nil {
			break
		}
		fmt.Printf("%s -> ", p.data)
		p = p.next
	}
	fmt.Println("End")
}

func (ll *LinkedList) AppendByValue(data interface{}) (bool, error) {
	if ll == nil {
		return false, errors.New("LinkedList is nil")
	}

	newNode := &Node{
		data: data,
		next: nil,
	}
	ll.Tail.next = newNode
	ll.Tail = newNode
	ll.length += 1

	return true, nil
}

func (ll *LinkedList) AppendByNode(data *Node) (bool, error) {
	if ll == nil {
		return false, errors.New("LinkedList is nil")
	}

	ll.Tail.next = data
	ll.Tail = data
	ll.length += 1

	return true, nil
}

func (ll *LinkedList) Insert(position int, data interface{}) (bool, error) {
	if ll == nil {
		return false, errors.New("LinkedList is nil")
	}

	newInsertNode := &Node{
		data: data,
		next: nil,
	}

	switch {
	case position <= 0:
		newInsertNode.next = ll.Head
		ll.Head = newInsertNode
		break
	case position >= ll.length:
		ll.AppendByNode(newInsertNode)
		break
	case position > 0 && position < ll.length:
		var cur = ll.Head
		var counter = 0
		for {
			if cur.next == nil || counter >= position-1 {
				break
			}
			cur = cur.next
			counter++
		}
		newInsertNode.next = cur.next
		cur.next = newInsertNode
		break
	default:
		break
	}

	return true, nil
}

func (ll *LinkedList) Delete(position int) (bool, error) {
	if ll == nil {
		return false, errors.New("LinkedList is nil")
	}

	switch {
	case position == 0:
		temp := ll.Head
		ll.Head = temp.next
		break;
	default:
		cur := ll.Head
		i := 0
		for {
			if cur.next == nil || i >= position - 1 {
				break
			}
			cur = cur.next
			i++
		}

		deletedNode := cur.next
		cur.next = deletedNode.next
		deletedNode = nil
	}
	return true, nil
}
