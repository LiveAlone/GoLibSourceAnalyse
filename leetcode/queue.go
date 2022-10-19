package leetcode

type Queue struct {
	pos  int
	data []interface{}
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) enqueue(e interface{}) {
	q.data = append(q.data, e)
}

func (q *Queue) pop() (rs interface{}) {
	rs = q.data[q.pos]
	q.data[q.pos] = nil
	q.pos += 1
	return
}

func (q *Queue) length() int {
	return len(q.data) - q.pos
}
