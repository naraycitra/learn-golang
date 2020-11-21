package main

import "fmt"

// Queue represent queue that hold slice of int
type Queue struct {
	datas []int
}

// Enqueue will add data at the end
func (q *Queue) Enqueue(i int) {
	q.datas = append(q.datas, i)
}

// Dequeue will remove the first data and return it
func (q *Queue) Dequeue() int {
	toRemove := q.datas[0]
	q.datas = q.datas[1:]
	return toRemove
}

func main() {
	var queue Queue
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	fmt.Println(queue)
	fmt.Println(queue.Dequeue())
	fmt.Println(queue)
}
