package list

import (
	"errors"
	"fmt"
)

type OneWayCircularLL struct {
	Head   *Node
	Tail   *Node
	Length int
}

func NewOneWayCirularLinkedList() *OneWayCircularLL {
	return &OneWayCircularLL{
		Head:   nil,
		Tail:   nil,
		Length: 0,
	}
}

func (cLL *OneWayCircularLL) Init(data []string) {
	for _, value := range data {

		if cLL.Head == nil && cLL.Tail == nil {
			firstNode := &Node{
				next: nil,
				data: value,
			}
			cLL.Head = firstNode
			cLL.Tail = firstNode
			cLL.Length += 1
			continue
		} else {
			cLL.AppendByValue(value)
		}
	}
}

func (cLL *OneWayCircularLL) AppendByValue(v interface{}) (bool, error) {
	if cLL == nil {
		return false, errors.New("Instance is nil")
	}

	if cLL.Head == nil && cLL.Tail == nil {
		firstNode := &Node{
			data: v,
			next: nil,
		}
		cLL.Head = firstNode
		cLL.Tail = firstNode
		cLL.Length += 1
		return true, nil
	}

	newNode := &Node{
		data: v,
		next: nil,
	}
	cLL.Tail.next = newNode
	cLL.Tail = newNode
	cLL.Tail.next = cLL.Head
	cLL.Length += 1

	return true, nil
}

func (cLL *OneWayCircularLL) Insert(position int, v string) (bool, error) {
	if cLL == nil {
		return false, errors.New("Instance is nil")
	}

	newNode := &Node {
		data: v,
		next: nil,
	}

	if position <= 0 {
		if cLL.Length == 1 {
			oldHead := cLL.Head
			cLL.Head = newNode
			cLL.Head.next = oldHead
			cLL.Tail = oldHead
			cLL.Tail.next = cLL.Head
			cLL.Length += 1
			return true, nil
		}

		newNode.next = cLL.Tail.next
		cLL.Tail.next = newNode
		cLL.Head = newNode
		cLL.Length += 1
		return true, nil
	}

	if position >= cLL.Length {
		if cLL.Length == 1 {
			cLL.Tail = newNode
			cLL.Head.next = cLL.Tail
			cLL.Tail.next = cLL.Head
			cLL.Length += 1
			return true, nil
		}

		cLL.Tail.next = newNode
		newNode.next = cLL.Head
		cLL.Tail = newNode
		cLL.Length += 1
		return true, nil
	}

	p := cLL.Head
	c := 0
	for {
		if c == position - 1 {
			break
		}

		c+=1
		p = p.next
	}

	newNode.next = p.next
	p.next = newNode
	cLL.Length += 1

	return true, nil
}

func (cLL *OneWayCircularLL) Delete(position int) (bool, error) {
	if cLL == nil {
		return false, errors.New("instance is nil")
	}
	if position < 0 || position >= cLL.Length {
		return false, errors.New("Invalid position")
	}

	switch {
	case cLL.Length == 0:
		return true, nil
	case cLL.Length == 1:
		cLL.Head = nil
		cLL.Tail = nil
		cLL.Length -= 1
		return true, nil
	case position == 0:
		newHead := cLL.Head.next
		if newHead == cLL.Tail {
			newHead.next = nil
			cLL.Tail.next = nil
			cLL.Head = newHead
			cLL.Tail = newHead
		} else {
			cLL.Head = newHead
			cLL.Tail.next = cLL.Head
		}

		cLL.Length -= 1
		return true, nil
	default:
		p:= cLL.Head
		c:= 0
		for {
			if p.next == nil || c >= position - 1 {
				break
			}
			p = p.next
			c++
		}

		// Special case
		if p.next == cLL.Tail {
			p.next = cLL.Tail.next
			cLL.Tail = p
			cLL.Length--
		} else {
			deletedNode := p.next
			p.next = deletedNode.next
			deletedNode.next = nil
			cLL.Length--
		}
	}

	return true, nil
}

func (cLL *OneWayCircularLL) OutputFromHead() {
	p := cLL.Head

	if cLL.Length == 0 {
		fmt.Println("Empty")
		return
	}

	if cLL.Length == 1 {
		fmt.Printf("%s -> Done", cLL.Head.data)
		fmt.Println()
		return
	}

	for {
		if p == nil {
			break
		}

		if p.next == cLL.Head {
			fmt.Printf("%s -> ", p.data)
			p = p.next
			fmt.Printf("%s", p.data)
			break
		}

		fmt.Printf("%s -> ", p.data)
		p = p.next
	}
	fmt.Println(" -> Done")
}

func (cLL *OneWayCircularLL) OutputFromTail() {
	p := cLL.Tail

	if cLL.Length == 0 {
		fmt.Println("Empty")
		return
	}

	if cLL.Length == 1 {
		fmt.Printf("%s -> Done", cLL.Tail.data)
		fmt.Println()
		return
	}

	for {
		if p == nil {
			return
		}

		if p.next == cLL.Tail {
			fmt.Printf("%s -> ", p.data)
			p = p.next
			fmt.Printf("%s", p.data)
			break
		}

		fmt.Printf("%s -> ", p.data)
		p = p.next
	}
	fmt.Println(" -> Done")
}

