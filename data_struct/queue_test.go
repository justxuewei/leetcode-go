package ds

import (
	"fmt"
	"testing"
)

func printQueueInt(queue *Queue) {
	if queue.head == queue.tail {
		return
	} else if queue.head < queue.tail {
		for i:=queue.head; i<queue.tail; i++ {
			intV, _ := queue.arr[i].(int)
			fmt.Printf("%d\t", intV)
		}
		fmt.Print()
		return
	} else {
		for i:=queue.head; i<len(queue.arr); i++ {
			intV, _ := queue.arr[i].(int)
			fmt.Printf("%d\t", intV)
		}
		for i:=0; i<queue.tail; i++ {
			intV, _ := queue.arr[i].(int)
			fmt.Printf("%d\t", intV)
		}
		fmt.Print()
	}
}

func TestQueue(t *testing.T) {
	queue := NewQueue(3)
	_ = queue.Enqueue(1)
	_ = queue.Enqueue(2)
	_ = queue.Enqueue(3)
	printQueueInt(queue)
}
