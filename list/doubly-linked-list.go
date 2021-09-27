package list

import (
	"errors"
	"fmt"
)

type DoublyLinkedList struct {
	Length int
	Head *DNode
	Tail *DNode
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		Length: 0,
		Head: nil,
		Tail: nil,
	}
}

func (dl *DoublyLinkedList) Init(data []string) error {
	if dl == nil {
		return errors.New("Doubly Linked-List is nil")
	}
	for _, value := range data {
		if dl.Head == nil && dl.Tail == nil {
			dNode := &DNode{
				data: value,
				prev: nil,
				next: nil,
			}
			dl.Head = dNode
			dl.Tail = dNode
			dl.Length = 1
		} else {
			dl.AppendByValue(value)
		}
	}
	return nil
}

func (dl *DoublyLinkedList) AppendByValue(data interface{}) error {
	if dl == nil {
		return errors.New("Doubly Linked-List is nil")
	}
	dNode := &DNode{
		data: data,
		prev: nil,
		next: nil,
	}
	if dl.Head == nil && dl.Tail == nil {
		dl.Head = dNode
		dl.Tail = dNode
		dl.Length = 1
		return nil
	}
	if dl.Length == 1 {
		dNode.prev = dl.Head
		dl.Head.next = dNode
		dl.Tail = dNode
		dl.Length += 1
		return nil
	}
	tail := dl.Tail
	dNode.prev = tail
	dl.Tail.next = dNode
	dl.Tail = dNode
	dl.Length+=1
	return nil
}


func (ld DoublyLinkedList) OutputFromHead() {
	h := ld.Head
	for {
		if h == nil {
			break
		}
		fmt.Printf("%s -> ", h.data)
		h = h.next
	}
	fmt.Println("End")
}
func (ld DoublyLinkedList) OutputFromTail() {
	t := ld.Tail
	for {
		if t == nil {
			break
		}
		fmt.Printf("%s -> ", t.data)
		t = t.prev
	}
	fmt.Println("Start")
}


