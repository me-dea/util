package pq

type PriorityQueue []*Node

type Node struct {
  Value string
  Priority int
}

func (pq PriorityQueue) Len() int {
  return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
  return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
  pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(n interface{}) {
  *pq = append(*pq, n.(*Node))
}

func (pq *PriorityQueue) Pop() interface{} {
  n := len(*pq)
  old := *pq
  el := old[n - 1]
  *pq = old[0: n -1]
  return el
}
