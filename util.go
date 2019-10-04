package gofunc

type Set struct {
	m map[interface{}]struct{}
}

func NewSet() *Set {
	m := make(map[interface{}]struct{})
	return &Set{m}
}

var Exist = struct{}{}

func (s *Set) Add(items ...interface{}) error {
	for i := range items {
		s.m[items[i]] = Exist
	}
	return nil
}

func (s *Set) Contains(item interface{}) bool {
	_, ok := s.m[item]
	return ok
}

func (s *Set) Size() int {
	return len(s.m)
}

func (s *Set) Clear() {
	s.m = make(map[interface{}]struct{})
}

func (s *Set) Delete(item interface{}) {
	delete(s.m, item)
}
