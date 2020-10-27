package queue


type Queue struct {
	queue []interface{}

	len int
}

func NewQueue() *Queue  {
	queue := &Queue{}
	queue.queue = make([]interface{}, 0)
	queue.len = 0

	return queue
}

func (q *Queue) Len() int {
	return q.len;
}

func (q *Queue) IsEmpty() bool {
	return q.len == 0;
}


