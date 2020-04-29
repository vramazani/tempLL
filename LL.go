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
	if list.len == 0 {
		list.first = node
		list.last = node
		list.len = 1
	} else {
		list.last.next = node
		node.prev = list.last
		list.last = node
		list.len++
	}
	return list.len
}

func (list *LinkedList) pop() (ret int, e error) {
	switch list.len {
	case 0:
		return 0, errors.New("empty")
	case 1:
		ret = list.first.data
		list.first = nil
		list.last = nil
		list.len = 0
		return
	case 2:
		ret = list.last.data
		list.first.next = nil
		list.last = list.first
		list.len = 1
		return
	default:
		ret = list.last.data
		list.last = list.last.prev
		list.last.next = nil
		list.len--
		return
	}
}

func (list *LinkedList) find(data, count int) (output []*Node) {
	if list.len == 0 || count < 1 {
		return
	} else {
		// output := []*Node {}
		current := list.first
		for i := 0; i < list.len; i++ {
			if current.data == data {
				output = append(output, current)
				if len(output) == count {
					return output
				}
			}
		}
		return output
	}
}

func (list *LinkedList) remove(data, count int) (countDone int) {
	nodes := list.find(data, count)
	for i := range nodes {
		nodes[i].prev.next = nodes[i].next
		nodes[i].next.prev = nodes[i].prev
		countDone++

	}
	list.len -= countDone
	return
}
