package main

//Queue object
type Queue struct {
	items []int
}

//Enqueue an item into the queue
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

//Dequeue an items into the queue
func (q *Queue) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}
