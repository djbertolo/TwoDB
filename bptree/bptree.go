package bptree

import "errors"

type Node struct {
	keys     []int
	values   []string
	parent   *Node
	children []*Node
	isLeaf   bool
}

type BPTree struct {
	order int
	root  *Node
}

// Node Functions
func NewNode(order int) *Node {
	return &Node{
		keys:     make([]int, order+1),
		values:   make([]string, order+1),
		children: make([]*Node, order+1),
	}
}

// Node Functions
func NewBPTree(order int) *BPTree {
	return &BPTree{
		order: order,
	}
}

func (b *BPTree) search(key int) (value string, err error) {
	var currNode *Node = b.root
	for {
		// Loop through keys in Node
		for i, v := range currNode.keys {

			// Check for conditions to go to left child
			if key < v && i == 0 && !currNode.isLeaf {
				currNode = currNode.children[0]
				break
			}

			// Check for conditions to go right child
			if (key == v || (i < len(currNode.keys)-1 && key < v)) && !currNode.isLeaf {
				currNode = currNode.children[i+1]
				break
			}

			// Conditions to break
			if key < v && i == 0 && currNode.isLeaf {
				value = currNode.values[i]
				return value, nil
			}

		}

		if currNode.isLeaf {
			return "", errors.New("Failed to get value with key.")
		}
	}
}

func (b *BPTree) getNode(key int) (*Node, error) {
	var currNode *Node = b.root
	for {
		// Loop through keys in Node
		for i, v := range currNode.keys {

			// Check for conditions to go to left child
			if key < v && i == 0 && !currNode.isLeaf {
				currNode = currNode.children[0]
				break
			}

			// Check for conditions to go right child
			if (key == v || (i < len(currNode.keys)-1 && key < v)) && !currNode.isLeaf {
				currNode = currNode.children[i+1]
				break
			}

			// Conditions to break
			if key < v && i == 0 && currNode.isLeaf {
				return currNode, nil
			}

			if key == v && currNode.isLeaf {
				return nil, errors.New("Duplicate Key found, returning nil")
			}

			if key > v && i == len(currNode.children)-1 && currNode.isLeaf {
				return currNode, nil
			}

		}
	}
}
