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

func printDLL(dll *DLL) {
	node := dll.head
	for node != nil {
		c := node.id + '0'
		if node.free {
			c = '.'
		}
		for i := 0; i < int(node.size); i++ {
			fmt.Print(string(c))
		}
		node = node.next
	}
	fmt.Println()
}

func main() {
	file, err := os.Open("data/d9/a.txt")
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
		if data[i] > 0 {
			pos = dll.tail.pos + int(dll.tail.size)
			newData = &Node{size: data[i], free: false, id: i, pos: pos, prev: dll.tail}
			dll.tail.next = newData
			dll.tail = newData
		}

		if i < len(free) && free[i] > 0 {
			pos = dll.tail.pos + int(dll.tail.size)
			newFree = &Node{size: free[i], free: true, id: i, pos: pos, prev: dll.tail}
			dll.tail.next = newFree
			dll.tail = newFree
		}
	}

	nodeToMove := dll.tail
	for nodeToMove != nil {
		if !nodeToMove.free {
			freeNode := dll.head
		movingLoop:
			for freeNode != nil && freeNode.pos < nodeToMove.pos {
				if freeNode.free && nodeToMove.size <= freeNode.size {
					movedSize := nodeToMove.size
					movedId := nodeToMove.id
					freeNode.size -= movedSize
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
						nodeToMove.next = next.next

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
						nodeToMove = prev
					}
					// Add new data
					newData = &Node{
						size: movedSize,
						free: false,
						id:   movedId,
						pos:  freeNode.pos,
						prev: freeNode.prev,
						next: freeNode,
					}
					if freeNode.prev != nil {
						freeNode.prev.next = newData
					}
					freeNode.prev = newData
					// Remove free space if filled, otherwise update position
					if freeNode.size == 0 {
						newData.next = freeNode.next
						if freeNode.next != nil {
							freeNode.next.prev = newData
						}
						freeNode.prev, freeNode.next = nil, nil
					} else {
						freeNode.pos += int(movedSize)
					}
					break movingLoop
				}
				freeNode = freeNode.next
			}
		}
		nodeToMove = nodeToMove.prev
		// printDLL(dll)
	}

	sum = 0
	node := dll.head
	for node != nil {
		if !node.free {
			s := int(node.size)
			sum += node.id * (s*node.pos + s*(s-1)/2)
		}
		node = node.next
	}
	fmt.Println("Whole move sum:", sum)
}
