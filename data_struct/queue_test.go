package ds

import (
	"fmt"
	"testing"
)

func printQueueInt(queue *Queue) {
	if queue.head == queue.tail {
		fmt.Print("Empty Queue")
	} else if queue.head < queue.tail {
		for i:=queue.head; i<queue.tail; i++ {
			intV, _ := queue.arr[i].(int)
			fmt.Printf("%d\t", intV)
		}
	} else {
		for i:=queue.head; i<len(queue.arr); i++ {
			intV, _ := queue.arr[i].(int)
			fmt.Printf("%d\t", intV)
		}
		for i:=0; i<queue.tail; i++ {
			intV, _ := queue.arr[i].(int)
			fmt.Printf("%d\t", intV)
		}
	}
	fmt.Println()
}

func TestQueue(t *testing.T) {
	queue := NewQueue(3)
	_ = queue.Enqueue(1)
	_ = queue.Enqueue(2)
	_ = queue.Enqueue(3)
	printQueueInt(queue)
	t.Log(queue.Len())
	first, _ := queue.Dequeue()
	firstInt, _ := first.(int)
	t.Logf("first: %d", firstInt)
	second, _ := queue.Dequeue()
	secondInt, _ := second.(int)
	t.Logf("second: %d", secondInt)
	printQueueInt(queue)
	t.Log(queue.Len())
	_ = queue.Enqueue(4)
	_ = queue.Enqueue(5)
	printQueueInt(queue)
	_, _ = queue.Dequeue()
	_, _ = queue.Dequeue()
	_, _ = queue.Dequeue()
	printQueueInt(queue)
	t.Log(queue.Len())
}
