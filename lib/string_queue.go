package lib

import "container/list"

type StringQueue interface {
	Clear()
	Pop() string
	Push(string)
	IsEmpty() bool
	Len() int
}

var _ StringQueue = &stringQueueImpl{}
var _ StringQueue = &uniqueStringQueueImpl{}

type stringQueueImpl struct {
	q *list.List
}

func NewStringQueue() *stringQueueImpl {
	return &stringQueueImpl{
		q: list.New(),
	}
}

func (q *stringQueueImpl) Clear() {
	q.q = list.New()
}

func (q *stringQueueImpl) Pop() string {
	e := q.q.Front()

	if e == nil {
		return ""
	}

	q.q.Remove(e)
	return e.Value.(string)
}

func (q *stringQueueImpl) Push(v string) {
	if v == "" {
		return
	}

	q.q.PushBack(v)
}

func (q *stringQueueImpl) IsEmpty() bool {
	return q.q.Len() == 0
}

func (q *stringQueueImpl) Len() int {
	return q.q.Len()
}

type uniqueStringQueueImpl struct {
	q *list.List
	strings map[string]struct{}
}

func NewUniqueStringQueue() *uniqueStringQueueImpl {
	return &uniqueStringQueueImpl{
		q: list.New(),
		strings : map[string]struct{}{},
	}
}

func (q *uniqueStringQueueImpl) Clear() {
	q.q = list.New()
	q.strings = map[string]struct{}{}
}

func (q *uniqueStringQueueImpl) Pop() string {
	e := q.q.Front()

	if e == nil {
		return ""
	}

	q.q.Remove(e)
	delete(q.strings, e.Value.(string))
	return e.Value.(string)
}

func (q *uniqueStringQueueImpl) Push(v string) {
	if v == "" {
		return
	}

	if _, ok := q.strings[v]; ok {
		return
	}
	q.q.PushBack(v)
	q.strings[v] = struct{}{}
}

func (q *uniqueStringQueueImpl) IsEmpty() bool {
	return q.q.Len() == 0
}

func (q *uniqueStringQueueImpl) Len() int {
	return q.q.Len()
}