package hw04_lru_cache //nolint:golint,stylecheck

type List interface {
	Len() int                          // длина списка
	Front() *listItem                  // первый элемент списка
	Back() *listItem                   // последний элемент списка
	PushFront(v interface{}) *listItem // добавить значение в начало
	PushBack(v interface{}) *listItem  // добавить значение в конец
	Remove(i *listItem)                // удалить элемент
	MoveToFront(i *listItem)           // переместить элемент в начало
}

type listItem struct {
	Value interface{} // значение
	Next  *listItem   // следующий элемент
	Prev  *listItem   // предыдущий элемент
}

type list struct {
	frontItem *listItem
	backItem  *listItem
	len       int
}

func (l list) Len() int {
	return l.len
}

func (l list) Front() *listItem {
	return l.frontItem
}

func (l list) Back() *listItem {
	return l.backItem
}

func (l *list) PushFront(v interface{}) *listItem {
	li := listItem{
		v,
		l.frontItem,
		nil,
	}

	if l.frontItem == nil {
		l.backItem = &li
	} else {
		l.frontItem.Prev = &li
	}

	l.frontItem = &li
	l.len++

	return &li
}

func (l *list) PushBack(v interface{}) *listItem {
	li := listItem{
		v,
		nil,
		l.backItem,
	}

	if l.backItem != nil {
		l.backItem.Next = &li
	} else {
		l.frontItem = &li
	}

	l.backItem = &li
	l.len++

	return &li
}

func (l *list) Remove(i *listItem) {
	// неплохо бы проверить, что элемент является членом этого списка, т.к. в задании сказано, что "Гарантируется, что методы Remove и MoveToFront вызываются от существующих в списке элементов." - пропускаем
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.frontItem = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.backItem = i.Prev
	}

	l.len--
}

func (l *list) MoveToFront(i *listItem) {
	// неплохо бы проверить, что элемент является членом этого списка, т.к. в задании сказано, что "Гарантируется, что методы Remove и MoveToFront вызываются от существующих в списке элементов." - пропускаем
	if i == l.frontItem {
		return
	}
	l.Remove(i)
	l.PushFront(i.Value)
}

func NewList() List {
	return &list{}
}
