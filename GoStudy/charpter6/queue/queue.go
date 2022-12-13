package queue

type Queue []int

func (q *Queue) Push(value int) {
	*q = append(*q, value)
}

func (q *Queue) Pop() (value int) {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}
