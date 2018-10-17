package datastructure

import (
	"sync"
)

// Double linked list implementation

type list struct {
	head   *Item
	tail   *Item
	len    int
	locker sync.RWMutex
}

type Item struct {
	Val  interface{}
	next *Item
	prev *Item
	list *list
}

func New() *list {
	return &list{}
}

func (list *list) Insert(value interface{}) {
	newItem := &Item{Val: value, list: list}
	list.locker.Lock()
	defer list.locker.Unlock()

	if list.head == nil {
		list.head = newItem
		list.tail = newItem
	} else {
		list.head.next = newItem
		newItem.prev = list.head
		list.head = newItem
	}

	list.len++
}

func (list *list) InsertAfterItem(value interface{}, item *Item) {
	if item == nil {
		list.Insert(value)
		return
	}
	newItem := &Item{Val: value, list: list}
	list.locker.Lock()
	defer list.locker.Unlock()

	newItem.prev = item.prev
	newItem.next = item
	if item.prev != nil { //if the item is not tail
		item.prev.next = newItem
		list.tail = newItem
	}
	item.prev = newItem

	list.len++
}

func (list *list) First() *Item {
	return list.head
}

func (list *list) Last() *Item {
	return list.tail
}

func (item *Item) Prev() *Item {
	return item.prev
}

func (item *Item) Next() *Item {
	return item.next
}

func (list *list) Has(value interface{}) bool {
	first := list.First()

	if first == nil {
		return false
	}

	for {
		if first.Val == value {
			return true
		}
		if first.next != nil {
			first = first.next
		} else {
			return false
		}
	}
}

func (list *list) Remove(val interface{}) {
	list.locker.RLock()
	first := list.First()
	last := list.Last()
	list.locker.RUnlock()

	list.locker.Lock()
	defer list.locker.Unlock()

	if first == last { // list has 1 element
		list.head = nil
		list.tail = nil
		list.len--
	} else if val == first.Val { // deleting head
		list.removeHead(first)
	} else { // deleting in the middle
		for {
			item := first.prev
			if item.Val == val { // val was found
				if item == last {
					list.removeTail(item)
				} else {
					list.removeMiddle(item)
				}
				return
			}
			if item == last { // val was not found
				return
			}
		}
	}
}

func (list *list) removeMiddle(item *Item) {
	item.prev.next = item.next.prev
	item.next.prev = item.prev.next
	list.len--
}

func (list *list) removeTail(last *Item) {
	last.next.prev = nil
	list.tail = last.next
	last = nil
	list.len--
}

func (list *list) removeHead(first *Item) {
	first.prev.next = nil
	list.head = first.prev
	first = nil
	list.len--
}

func (list *list) Length() int {
	return list.len
}
