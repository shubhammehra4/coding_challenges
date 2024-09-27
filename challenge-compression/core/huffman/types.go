package huffman

type Node struct {
	Value  rune
	Count  int
	Left   *Node
	Right  *Node
	Parent *Node
}

type PrefixCode struct {
	Code uint64
	Bits int32
}

func newLeafNode(value rune, weight int) *Node {
	return &Node{Value: value, Count: weight}
}

func newInternalNode(left, right *Node) *Node {
	return &Node{Count: left.Count + right.Count, Left: left, Right: right}
}

// implements container/heap.Interface
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].Count == pq[j].Count {
		return pq[i].Value < pq[j].Value
	}
	return pq[i].Count < pq[j].Count
}
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*Node)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}
