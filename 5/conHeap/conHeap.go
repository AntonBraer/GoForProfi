package main

import (
	"container/heap"
	"fmt"
)

type heapFloat32 []float32

func (n *heapFloat32) Pop() interface{} {
	old := *n
	x := old[len(old)-1]
	new := old[0 : len(old)-1]
	*n = new
	return x
}

func (n *heapFloat32) Push(x interface{}){
	*n = append(*n, x.(float32))
}

func (n heapFloat32) Len() int {
	return len(n)
}

func (n heapFloat32) Less(a, b int) bool {
	return n[a] < n[b]
}

func (n heapFloat32) Swap(a, b int) {
	n[a], n[b] = n[b], n[a]
}

func main() {
	myHeap := &heapFloat32{1.2, 2.1, 3.1, 4.5}
	heap.Init(myHeap)
	size := len(*myHeap)
	fmt.Println("size:", size)
	fmt.Printf("%v\n", myHeap)

	myHeap.Push(float32(5.5))
	myHeap.Push(float32(0.2))
	fmt.Println("size:", len(*myHeap))
	fmt.Printf("%v\n", myHeap)
	tmp := myHeap.Pop()
	tmp = myHeap.Pop()
	fmt.Println("tmp:", tmp)
	fmt.Printf("%v\n", myHeap)
}