package PriorityQueue

import (
  "container/heap"
  "testing"
  "fmt"
)

func TestSimple(t *testing.T) {
  var pq PriorityQueue
  heap.Push(&pq, &Node{
    Value: "s",
    Priority: 1,
  })
  el := heap.Pop(&pq).(*Node)
  expected := Node{
    Value: "s",
    Priority: 1,
  }
  if el.Priority != 1 || el.Value != "s" {
    fmt.Printf("unexpected value %+v, expected %+v\n", *el, expected)
  }
}
