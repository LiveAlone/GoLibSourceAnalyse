package leetcode

type Stack struct {
	pos  int
	data []interface{}
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) push(e interface{}) {
	if len(s.data) <= s.pos {
		s.data = append(s.data, e)
		s.pos = s.pos + 1
		return
	}
	s.data[s.pos] = e
	s.pos = s.pos + 1
}

func (s *Stack) pop() interface{} {
	rs := s.data[s.pos-1]
	s.data[s.pos-1] = nil
	s.pos = s.pos - 1
	return rs
}

func (s *Stack) top() interface{} {
	return s.data[s.pos]
}

func (s *Stack) length() int {
	return s.pos
}

func (s *Stack) Data() []interface{} {
	return s.data
}
