package utils
import "container/list"

// Queue is a queue
type Queue interface {
	Front() *list.Element
	Len() int
	Add(interface{})
	Remove() *list.Element
}
//QueueImpl
type QueueImpl struct {
	*list.List
}

func (q *QueueImpl) Add(v interface{}) {
	q.PushBack(v)
}
func (q *QueueImpl) Len() int{
	return q.List.Len()
}
func (q *QueueImpl) Remove() *list.Element {
	e := q.Front()
	q.List.Remove(e)
	return e
}

// New is a new instance of a Queue
func New() Queue {
	return &QueueImpl{list.New()}
}