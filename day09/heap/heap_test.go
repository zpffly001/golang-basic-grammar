package heap_test

import (
	"fmt"
	"testing"

	"gitee.com/infraboard/go-course/day09/heap"
)

func TestHeap(t *testing.T) {
	m := []int{0, 9, 3, 6, 2, 1, 7} //第0个下标不放目标元素
	h := heap.NewIntHeap(m)
	fmt.Println(h.Items())

	// Push()：插入元素
	h.Push(50)
	fmt.Println(h.Items())

	// Pop()：移除并返回堆顶元素(这个好像默认是大顶堆，每次移除的都是最大的)
	h.Pop()
	fmt.Println(h.Items())

	h.Pop()
	fmt.Println(h.Items())

	h.Pop()
	fmt.Println(h.Items())
	h.Pop()
	fmt.Println(h.Items())
}

func TestBuildHeap(t *testing.T) {
	heap.Example_intHeap()
}

func TestPriorityQueue(t *testing.T) {
	// 测试优先队列
	heap.TestPriorityQueue()
}
