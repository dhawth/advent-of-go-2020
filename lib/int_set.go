package lib

import "encoding/json"

type IntSet interface {
	Add(int)
	Contains(int) bool
	IsEmpty() bool
	Len() int
	Members() []int
}

var _ StringSet = &stringSetImpl{}
var _ json.Marshaler = &stringSetImpl{}

type intSetImpl struct {
	members map[int]struct{}
}

func NewIntSet() *intSetImpl {
	return &intSetImpl{
		members: map[int]struct{}{},
	}
}

func (s *intSetImpl) Add(v int) {
	s.members[v] = struct{}{}
}

func (s *intSetImpl) Contains(v int) bool {
	_, ok := s.members[v]
	return ok
}

func (s *intSetImpl) IsEmpty() bool {
	return len(s.members) == 0
}

func (s *intSetImpl) Len() int {
	return len(s.members)
}

func (s *intSetImpl) Members() []int {
	var res []int
	for k := range s.members {
		res = append(res, k)
	}
	return res
}

func (s *intSetImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Members())
}
