package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	size byte
	free bool
	id   int
	pos  int
	prev *Node
	next *Node
}

type DLL struct {
	head *Node
	tail *Node
}

func main() {
	file, err := os.Open("data/d9/test.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		fmt.Println("Empty file")
		return
	}
	line := scanner.Text()

	var data, free []byte
	for i, c := range line {
		if i%2 == 0 {
			data = append(data, byte(c)-'0')
		} else {
			free = append(free, byte(c)-'0')
		}
	}
	dataOriginal := make([]byte, len(data))
	copy(dataOriginal, data)
	freeOriginal := make([]byte, len(free))
	copy(freeOriginal, free)

	sum, pos := 0, 0
outer:
	for i, j := 0, len(data)-1; i < len(data); i++ {
		d := int(data[i])
		sum += i * (d*pos + d*(d-1)/2)
		pos += d

		for free[i] > 0 {
			if j <= i {
				break outer
			}

			moved := min(free[i], data[j])
			free[i] -= moved
			data[j] -= moved

			m := int(moved)
			sum += j * (m*pos + m*(m-1)/2)
			pos += m

			if data[j] == 0 {
				j--
			}
		}
	}
	fmt.Println("Partial move sum:", sum)

	// Part 2
	copy(data, dataOriginal)
	copy(free, freeOriginal)

	dll := &DLL{}
	newData := &Node{size: data[0], free: false, id: 0, pos: 0}
	dll.head = newData
	dll.tail = newData

	newFree := &Node{size: free[0], free: true, prev: dll.head, id: 0, pos: int(data[0])}
	dll.tail = newFree
	dll.head.next = newFree

	for i := 1; i < len(data); i++ {
		pos = dll.tail.pos + int(dll.tail.size)
		newData = &Node{size: data[i], free: false, id: i, pos: pos, prev: dll.tail}
		dll.tail.next = newData
		dll.tail = newData

		if i < len(free) {
			pos = dll.tail.pos + int(dll.tail.size)
			newFree = &Node{size: free[i], free: true, id: i, pos: pos, prev: dll.tail}
			dll.tail.next = newFree
			dll.tail = newFree
		}
	}

	node := dll.head
	for node != nil {
		if node.free {
			nodeToMove := dll.tail
		movingLoop:
			for nodeToMove != nil {
				if !nodeToMove.free && nodeToMove.size <= node.size {
					node.size -= nodeToMove.size
					nodeToMove.free = true
					// Merge free space after
					next := nodeToMove.next
					if next != nil && next.free {
						nodeToMove.size += next.size
						if next.next != nil {
							next.next.prev = nodeToMove
						} else {
							dll.tail = nodeToMove
						}
						next = next.next

						next.prev, next.next = nil, nil
					}
					// Merge free space before
					prev := nodeToMove.prev
					if prev != nil && prev.free {
						prev.size += nodeToMove.size
						prev.next = nodeToMove.next
						if nodeToMove.next != nil {
							nodeToMove.next.prev = prev
						} else {
							dll.tail = prev
						}

						nodeToMove.prev, nodeToMove.next = nil, nil
					}
					// Add new data
					newData = &Node{
						size: nodeToMove.size,
						free: false,
						id:   nodeToMove.id,
						pos:  node.pos,
						prev: node.prev,
						next: node,
					}
					if node.prev != nil {
						node.prev.next = newData
					}
					node.prev = newData
					// Remove free space if filled, otherwise update position
					if node.size == 0 {
						node.prev.next = node.next
						node.prev, node.next = nil, nil
						break movingLoop
					} else {
						node.pos += int(nodeToMove.size)
					}
				}
				nodeToMove = nodeToMove.prev
			}
		}
		node = node.next
	}

	sum = 0
	node = dll.head
	for node != nil {
		fmt.Println(node)
		node = node.next
	}
}
