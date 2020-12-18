package lib

import "encoding/json"

type StringSet interface {
	Add(string)
	Contains(string) bool
	IsEmpty() bool
	Len() int
	Members() []string
}

var _ StringSet = &stringSetImpl{}
var _ json.Marshaler = &stringSetImpl{}

type stringSetImpl struct {
	members map[string]struct{}
}

func NewStringSet() *stringSetImpl {
	return &stringSetImpl{
		members: map[string]struct{}{},
	}
}

func (s *stringSetImpl) Add(v string) {
	if v == "" {
		return
	}
	s.members[v] = struct{}{}
}

func (s *stringSetImpl) Contains(v string) bool {
	if v == "" {
		return false
	}
	_, ok := s.members[v]
	return ok
}

func (s *stringSetImpl) IsEmpty() bool {
	return len(s.members) == 0
}

func (s *stringSetImpl) Len() int {
	return len(s.members)
}

func (s *stringSetImpl) Members() []string {
	var res []string
	for k := range s.members {
		res = append(res, k)
	}
	return res
}

func (s *stringSetImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Members())
}

