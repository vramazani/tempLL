package main

import "errors"

func main() {

}

type Node struct {
	data int
	prev *Node
	next *Node
}

type LinkedList struct {
	first *Node
	last  *Node
	len   int
}

func (list *LinkedList) append(node *Node) int {
	// consider both empty and non-empty list
	if 0 == list.len {
		list.first = node
	} else {
		list.last.next = node
		node.prev = list.last
	}
	list.last = node
	list.len++
	return list.len
}

func (list *LinkedList) pop() (ret int, e error) {
	if 0 == list.len {
		return 0, errors.New("empty")
	}
	ret = list.last.data
	if 1 == list.len {
		list.first = nil
		list.last = nil
	} else {
		list.last = list.last.prev
		list.last.next = nil
	}
	list.len--
	return
}

func (list *LinkedList) find(data, count int) (output []*Node) {
	// def: finds [first] up to count occurances of data in list
	if 0 == list.len || 0 == count {
		return
	} else {
		current := list.first
		for i := 0; i < list.len; i++ {
			if current.data == data {
				output = append(output, current)
				if len(output) == count {
					return
				}
			}
		}
		return
	}
}

func (list *LinkedList) remove(data, count int) (countDone int) {
	// def: removes [first] up to count occurances of data in list
	nodes := list.find(data, count)

	// list.len in the for : 0, 1, 2, 3+
	// 0: only when initially empty; skips for cuz LL.find yielded an empty array
	// 1: first&&last
	// 2: first / last
	// 3+: first / last / inner
	for i := range nodes {
		if nodes[i] == list.first && nodes[i] == list.last {
			// here list.len == 1
			// hence must be the last iteration
			list.first = nil
			list.last = nil
		} else if nodes[i] == list.first {
			// here list.len >= 2
			list.first = list.first.next
			list.first.prev = nil
		} else if nodes[i] == list.last {
			// here list.len >= 2
			list.last = list.last.prev
			list.last.next = nil
		} else { // inner
			// here list.len >= 3
			nodes[i].prev.next = nodes[i].next
			nodes[i].next.prev = nodes[i].prev
		}
		list.len--
		countDone++
	}
	// here countDone == len(nodes)
	return
}
