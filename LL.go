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

func (list *LinkedList) put(node *Node, index int) (int, error) {
	// (push_before) def: put node into list so that its position becomes index
	if index < 0 || list.len < index {
		return -1, errors.New("illegal index")
	}

	// by list.len 0, 1, 2, 3+
	// 0: case 0 if
	// 1: case 0 else /
	switch index {
	case 0: // becomes first
		if list.len == 0 {
			list.append(node)
		} else {
			node.next = list.first
			node.next.prev = node
			list.first = node
		}
	// case 1:
	// case list.len - 1:
	case list.len: // becomes last
		list.append(node)
	default: // becomes an inner
		toBeNext := list.first
		for i := 0; i < list.len; i++ {
			if i == index {
				break
			}
			toBeNext = toBeNext.next
		}
		node.next = toBeNext
		node.prev = toBeNext.prev
		node.prev.next = node
		node.next.prev = node
	}
	return list.len, nil
}

func (list *LinkedList) find(data, count int) (output []*Node) {
	// def: forward finds up to count occurances of data in list
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
	// def: forward removes up to count occurances of data in list
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
			// same effect as list.pop()
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
