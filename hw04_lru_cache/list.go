package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Prev  *ListItem
	Next  *ListItem
	Key   Key
}

type list struct {
	first *ListItem
	last  *ListItem
	len   int
	items []ListItem
}

func (list *list) Len() int {
	return list.len
}

func (list *list) Front() *ListItem {
	return list.first
}

func (list *list) Back() *ListItem {
	return list.last
}

func (list *list) PushFront(v interface{}) *ListItem {
	item := &ListItem{
		Value: v,
		Next:  list.first,
		Prev:  nil,
	}

	if list.first != nil {
		list.first.Prev = item
	}

	list.first = item
	if list.last == nil {
		list.last = item
	}

	list.len++
	return item
}

func (list *list) PushBack(v interface{}) *ListItem {
	item := &ListItem{
		Value: v,
		Next:  nil,
		Prev:  list.last,
	}

	if list.last != nil {
		list.last.Next = item
	}

	list.last = item
	if list.first == nil {
		list.first = item
	}

	list.len++
	return item
}

func (list *list) Remove(item *ListItem) {
	_ = list.moveItem(item)
	list.len--
}

func (list *list) moveItem(item *ListItem) *ListItem {
	switch {
	case list.len == 1:
		item.Prev = nil
		item.Next = nil
	case item.Next == nil:
		item.Prev.Next = nil
		list.last = item.Prev
	case item.Prev == nil:
		item.Next.Prev = nil
		list.first = item.Next
	default:
		item.Next.Prev = item.Prev
		item.Prev.Next = item.Next
	}

	return item
}

func (list *list) MoveToFront(item *ListItem) {
	if list.Front() == item {
		return
	}

	list.PushFront(item.Value)
	list.Remove(item)
}

func NewList() List {
	return new(list)
}

func (list *list) GetItems() []ListItem {
	return list.items
}
