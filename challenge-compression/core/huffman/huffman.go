package huffman

import (
	"container/heap"

	"github.com/samber/lo"
)

// Build builds a Huffman tree from a frequency map.
func Build(freqMap map[rune]int) *Node {
	leaves := lo.Map(lo.Entries(freqMap), func(k lo.Entry[rune, int], _ int) *Node { return newLeafNode(k.Key, k.Value) })
	pq := PriorityQueue(leaves)
	heap.Init(&pq)
	return buildSorted(PriorityQueue(leaves))
}

// buildSorted builds a Huffman tree from a priority queue of nodes.
func buildSorted(pq PriorityQueue) *Node {
	if len(pq) == 0 {
		return nil
	}
	for pq.Len() > 1 {
		left, right := heap.Pop(&pq).(*Node), heap.Pop(&pq).(*Node)
		parentNode := newInternalNode(left, right)
		heap.Push(&pq, parentNode)
	}
	return heap.Pop(&pq).(*Node)
}

// PrefixCodeTableOld builds a table of prefix codes from a Huffman tree.
func PrefixCodeTableOld(root *Node) map[rune]string {
	table := make(map[rune]string)
	var traverse func(n *Node, code string)
	traverse = func(n *Node, code string) {
		if n.Left == nil && n.Right == nil {
			table[n.Value] = code
			return
		}
		traverse(n.Left, code+"0")
		traverse(n.Right, code+"1")
	}

	traverse(root, "")
	return table
}

func PrefixCodeTable(root *Node) map[rune]PrefixCode {
	table := make(map[rune]PrefixCode)

	var traverse func(n *Node, code uint64, bits int32)
	traverse = func(n *Node, code uint64, bits int32) {
		if n.Left == nil {
			table[n.Value] = PrefixCode{Code: code, Bits: bits}
			return
		}
		bits++
		traverse(n.Left, code<<1, bits)
		traverse(n.Right, code<<1+1, bits)
	}

	traverse(root, 0, 0)
	return table
}
