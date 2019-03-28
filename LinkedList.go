package main

import (
	"errors"
	"fmt"
)

type Node struct {
	data int
	link *Node
}

type LinkedList struct {
	head  *Node
	_size int
}

func (list *LinkedList) add(data int) {
	defer list.increaseSize()
	newNode := Node{data: data, link: nil}
	if list.head == nil {
		list.head = &Node{data: -1, link: &newNode}

	} else {
		pointer := list.head.link
		for ; pointer.link != nil; {
			pointer = pointer.link
		}
		pointer.link = &newNode
	}
}

func (list *LinkedList) increaseSize() {
	list._size += 1
}

func (list *LinkedList) decreaseSize() {
	list._size -= 1
}

func (list *LinkedList) size() int {
	return list._size
}

func (list *LinkedList) get(i int) (error, int) {
	if i < 0 || i >= list.size() {
		return errors.New("out of bounds"), -1
	}
	counter := 0
	pointer := list.head.link
	for ; pointer != nil; {
		if counter == i {
			return nil, pointer.data
		}
		counter ++
		pointer = pointer.link
	}
	return errors.New("out of bounds"), -1
}

func (list *LinkedList) search(target int) (error, int) {
	pointer := list.head.link
	counter := 0
	for ; pointer != nil; {
		if pointer.data == target {
			return nil, counter
		}
		counter++
		pointer = pointer.link
	}
	return errors.New("element not found"), -1
}

func (list *LinkedList) getAll() []int {
	var returnArray []int
	if list.size() == 0 {
		return returnArray
	}

	pointer := list.head.link
	for ; pointer != nil; {
		returnArray = append(returnArray, pointer.data)
		pointer = pointer.link
	}

	return returnArray
}

func (list *LinkedList) removeAtIndex(i int) {
	if i < 0 || i >= list.size() {
		fmt.Println(errors.New(fmt.Sprintf("index %d out of bounds", i)))
		return
	}
	counter := 0
	pointer := list.getHeadPointer()
	frontPointer := pointer.link
	for ; frontPointer != nil; {
		if counter == i {
			pointer.link = frontPointer.link
			list.decreaseSize()
			break
		}

		counter++
		pointer = pointer.link
		frontPointer = frontPointer.link
	}
}

func (list *LinkedList) getHeadPointer() *Node {
	return list.head
}

func main() {

	var linkedList LinkedList

	linkedList.add(1)
	linkedList.add(2)
	linkedList.add(3)

	fmt.Println(linkedList.get(0))
	fmt.Println(linkedList.get(1))
	fmt.Println(linkedList.get(2))

	fmt.Println()

	fmt.Println(linkedList.search(1))
	fmt.Println(linkedList.search(2))
	fmt.Println(linkedList.search(3))

	fmt.Println()
	fmt.Println(linkedList.getAll())

	linkedList.removeAtIndex(2)
	fmt.Println(linkedList.getAll())

	_, index := linkedList.search(1)
	linkedList.removeAtIndex(index)
	fmt.Println(linkedList.getAll())

}
