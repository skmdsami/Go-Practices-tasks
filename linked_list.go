package main

import (
	"fmt"
	"errors"
)

type NumericTypes interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 |  float32 | float64
}

type Node[T NumericTypes] struct {
	value T
	next *Node[T]
}

type linkedList[T NumericTypes] struct{
	head *Node[T]
	tail *Node[T]
}


func node[T NumericTypes](val T) *Node[T] {
	return &Node[T]{
		value: val,
		next:  nil,
	}
}

func LinkedList[T NumericTypes]()  *linkedList[T] {
	return &linkedList[T]{
		head: nil,
		tail: nil,
	}
}


func (l *linkedList[T]) add(val T) {
	if l.head == nil {
		l.head = node(val)
		l.tail = l.head
	}else{
		l.tail.next = node(val)
		l.tail = l.tail.next
	}
}


func (l *linkedList[T]) traverse() {
	val := l.head ;
	for val != nil {
		fmt.Println(val.value)
		val = val.next
	}
}


func (l *linkedList[T]) contains(value T) (bool, error) {
		val := l.head
		for val != nil {
			if val.value == value {
				return true, nil
			}
			val = val.next
		}
		return false, errors.New("element not found")
}

func main() {
	fmt.Println("Creating a linked list")
	linked_list := LinkedList[uint]()
	linked_list.add(5)
	linked_list.add(4)
	linked_list.add(3)
	linked_list.add(2)
	linked_list.add(1)
	linked_list.traverse()
	fmt.Println("let's traverse again")
	linked_list.traverse()
	val1,val2 := linked_list.contains(7)
	fmt.Printf("contains : %b error : %s \n",val1,val2)
	val := fmt.Errorf("cannot divide by zero")
	fmt.Println("error was :",val)
	panic("i only did this for fun really no error occured")
}

// there is no exceptions to throw on , in go actually so there is no need of try catch blocks
// go panics for some cases where , when dealing with array index out of bounds, Nil pointer dereferences for those conditions