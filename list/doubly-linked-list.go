package list

import (
	"errors"
	"fmt"
)

type DoublyLinkedList struct {
	Length int
	Head   *DNode
	Tail   *DNode
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		Length: 0,
		Head:   nil,
		Tail:   nil,
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
	switch {
	case dl.Head == nil && dl.Tail == nil:
		dl.Head = dNode
		dl.Tail = dNode
		dl.Length = 1
		break
	case dl.Length == 1:
		dNode.prev = dl.Head
		dl.Head.next = dNode
		dl.Tail = dNode
		dl.Length += 1
		break
	default:
		tail := dl.Tail
		dNode.prev = tail
		dl.Tail.next = dNode
		dl.Tail = dNode
		dl.Length += 1
		break
	}
	return nil
}
func (dl *DoublyLinkedList) Insert(v interface{}, position int) error {
	if dl == nil {
		return errors.New("Doubly Linked-list is nil")
	}
	dnode := &DNode{
		data: v,
		next: nil,
		prev: nil,
	}
	switch {
	case position <= 0:
		dnode.next = dl.Head
		dl.Head.prev = dnode
		dl.Head = dnode
		dl.Length += 1
		break
	case position >= dl.Length:
		dl.AppendByValue(v)
		break
	default:
		i := 0
		curr := dl.Head
		for {
			if i >= position-1 || curr.next == nil {
				break
			}
			curr = curr.next
			i++
		}
		dnode.next = curr.next
		dnode.prev = curr
		curr.next.prev = dnode
		curr.next = dnode
		dl.Length += 1
		break
	}
	return nil
}
func (dl *DoublyLinkedList) Delete(position int) error{
	if dl == nil {
		return errors.New("Doubly Linked-list is nil")
	}
	switch {
	case position <= 0:
		oldHead := dl.Head
		dl.Head = dl.Head.next
		dl.Head.prev = nil
		oldHead.next = nil
		dl.Length -= 1
		break;
	default:
		i := 0
		curr := dl.Head
		for {
			if i >= position - 1 || curr.next == nil {
				break
			}
			curr = curr.next
			i += 1
		}
		if curr == dl.Tail {
			oldTail := curr
			dl.Tail = dl.Tail.prev
			dl.Tail.next = nil
			oldTail.next = nil
			oldTail.prev = nil
			oldTail = nil
			dl.Length -= 1
			break;
		}

		oldNode := curr.next
		curr.next = oldNode.next
		oldNode.next.prev = curr
		oldNode.next = nil
		oldNode.prev = nil
		oldNode = nil
		dl.Length -= 1
		break
	}
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
