package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (list *linkedList) prepend(node *node) {
	second := list.head
	list.head = node
	list.head.next = second
	list.length++
}

func (list linkedList) printData() {
	toPrint := list.head
	for list.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		list.length--
	}
	fmt.Println()
}

func (list *linkedList) deleteWithValue(v int) {
	// handle empty list
	if list.length == 0 {
		return
	}

	// handle remove head value
	if list.head.data == v {
		list.head = list.head.next
		list.length--
		return
	}

	previousToDelete := list.head

	// handle if only one element in the linked list
	if previousToDelete.next == nil {
		if previousToDelete.data == v {
			list.head = &node{}
			list.length--
		}
		return
	}

	for previousToDelete.next.data != v {
		// handle value not exist
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	list.length--
}

func main() {
	node1 := &node{data: 12}
	node2 := &node{data: 24}
	node3 := &node{data: 36}
	node4 := &node{data: 48}
	var list linkedList
	list.prepend(node1)
	list.prepend(node2)
	list.prepend(node3)
	list.prepend(node4)
	list.printData()
	list.deleteWithValue(36)
	list.printData()
	list.deleteWithValue(48)
	list.printData()
	list.deleteWithValue(12)
	list.printData()
	list.deleteWithValue(24)
	list.printData()
	node5 := &node{data: 64}
	list.prepend(node5)
	list.printData()
	list.deleteWithValue(100)
	var list2 linkedList
	list2.deleteWithValue(1)
}
