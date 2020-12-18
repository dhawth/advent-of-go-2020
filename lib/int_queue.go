package lib

import (
	"container/list"
	"encoding/json"
	"fmt"
)

var (
	ErrQueueEmpty = fmt.Errorf("queue is empty")
)

type IntQueue interface {
	Clear()
	// Pop pops from the front (FIFO)
	Pop() (int, error)
	// Push pushes items to the rear (FIFO)
	Push(int)
	// Shift pops from the rear (LIFO)
	Shift() (int, error)
	// UnShift pushes items to the front (LIFO)
	UnShift(int)
	IsEmpty() bool
	Len() int
}

var _ IntQueue = &intQueueImpl{}
var _ IntQueue = &uniqueIntQueueImpl{}
var _ json.Marshaler = &intQueueImpl{}
var _ json.Marshaler = &uniqueIntQueueImpl{}

type intQueueImpl struct {
	q *list.List
}

func NewIntQueue() *intQueueImpl {
	return &intQueueImpl{
		q: list.New(),
	}
}

func (q *intQueueImpl) Clear() {
	q.q = list.New()
}

// Pop pops and returns values from the front of the queue
func (q *intQueueImpl) Pop() (int, error) {
	e := q.q.Front()

	if e == nil {
		return 0, ErrQueueEmpty
	}

	q.q.Remove(e)
	return e.Value.(int), nil
}

func (q *intQueueImpl) Push(v int) {
	q.q.PushBack(v)
}

// Shift pops and returns values from the end of the queue
func (q *intQueueImpl) Shift() (int, error) {
	e := q.q.Back()
	if e == nil {
		return 0, ErrQueueEmpty
	}

	q.q.Remove(e)
	return e.Value.(int), nil
}

func (q *intQueueImpl) UnShift(v int) {
	q.q.PushFront(v)
}

func (q *intQueueImpl) IsEmpty() bool {
	return q.q.Len() == 0
}

func (q *intQueueImpl) Len() int {
	return q.q.Len()
}

func (q *intQueueImpl) MarshalJSON() ([]byte, error) {
	elements := make([]int, q.q.Len())

	index := 0
	for item := q.q.Front(); item != nil; item = item.Next() {
		elements[index] = item.Value.(int)
		index++
	}

	return json.Marshal(elements)
}

type uniqueIntQueueImpl struct {
	q *list.List
	ints map[int]struct{}
}

func NewUniqueIntQueue() *uniqueIntQueueImpl {
	return &uniqueIntQueueImpl{
		q: list.New(),
		ints : map[int]struct{}{},
	}
}

func (q *uniqueIntQueueImpl) Clear() {
	q.q = list.New()
	q.ints = map[int]struct{}{}
}

func (q *uniqueIntQueueImpl) Pop() (int, error) {
	e := q.q.Front()

	if e == nil {
		return 0, ErrQueueEmpty
	}

	q.q.Remove(e)
	delete(q.ints, e.Value.(int))
	return e.Value.(int), nil
}

func (q *uniqueIntQueueImpl) Push(v int) {
	if _, ok := q.ints[v]; ok {
		return
	}
	q.q.PushBack(v)
	q.ints[v] = struct{}{}
}

func (q *uniqueIntQueueImpl) Shift() (int, error) {
	e := q.q.Back()

	if e == nil {
		return 0, ErrQueueEmpty
	}

	q.q.Remove(e)
	delete(q.ints, e.Value.(int))
	return e.Value.(int), nil
}

func (q *uniqueIntQueueImpl) UnShift(v int) {
	if _, ok := q.ints[v]; ok {
		return
	}
	q.q.PushFront(v)
	q.ints[v] = struct{}{}
}

func (q *uniqueIntQueueImpl) IsEmpty() bool {
	return q.q.Len() == 0
}

func (q *uniqueIntQueueImpl) Len() int {
	return q.q.Len()
}

func (q *uniqueIntQueueImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(q.q)
}