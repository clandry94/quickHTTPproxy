package queue

type Queue []string

func (q *Queue) Push(s string) {
	*q = append(*q, s)
}

func (q *Queue) Pop() (s string) {
	s = (*q)[0]
	*q = (*q)[1:]
	return
}

func (q *Queue) Len() int {
	return len(*q)
}
