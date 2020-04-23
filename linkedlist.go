package main

import "fmt"

type Dictionary struct {
	name  string
	value interface{}
	next  *Dictionary
	prev  *Dictionary
}

type List struct {
	head *Dictionary
	tail *Dictionary
}

func (l *List) addItem(name string, value interface{}) {
	dictionary := &Dictionary{
		name:  name,
		value: value,
	}
	if l.head == nil {
		l.head = dictionary
	} else {
		l.tail.next = dictionary
		dictionary.prev = l.tail
	}
	l.tail = dictionary
}

func (l *List) reverse() {
	item := l.head
	var prev *Dictionary
	for item != nil {
		temp := item.next
		item.next = prev
		prev = item
		item = temp
	}
	l.head = prev
}

func (l *List) display() {
	item := l.head
	for item != nil {
		fmt.Println(item.value)
		item = item.next
	}
}

func (l *List) displayBackward() {
	item := l.tail
	for item != nil {
		fmt.Println(item.value)
		item = item.prev
	}
}

func testLinkedList() {
	list := List{}
	list.addItem("1", "1")
	list.addItem("2", "2")
	list.addItem("3", "3")
	list.addItem("4", "4")

	list.display()
	list.reverse()
	list.display()
}
