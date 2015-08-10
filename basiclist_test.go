package skiplist

import (
	"fmt"
	"testing"
)

func TestLinkedNode(t *testing.T) {
	headNode := basicnote{key: 0, nextNode: nil}
	headNode.nextNode = &basicnote{key: 1, val: "node1", nextNode: nil}
	headNode.nextNode.nextNode = &basicnote{key: 2, val: "node2", nextNode: nil}

	fmt.Printf("head->")
	currentNode := &headNode
	for currentNode.nextNode != nil {
		fmt.Printf("[key:%d][val:%v]->", currentNode.key, currentNode.val)
		currentNode = currentNode.nextNode
	}
	fmt.Printf("[key:%d][val:%v]->", currentNode.key, currentNode.val)
	fmt.Printf("nil\n")
}

func TestBasicList(t *testing.T) {
	bList := NewBasicList()
	bList.DisplayAll()
	bList.Insert(3, "string3")
	bList.DisplayAll()
}
