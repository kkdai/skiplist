package skiplist

import (
	"errors"
	"fmt"
)

type skipnode struct {
	key     int
	val     interface{}
	forward []*skipnode
}

type skiplist struct {
	header *skipnode
}

//NewSkipList : Init structure for basic Sorted Linked List.
func NewSkipList() *skiplist {
	var empty interface{}
	return &skiplist{header: &skipnode{key: 0, val: empty, forward: nil}}
}

func (b *skiplist) Insert(key int, value interface{}) {
}

func (b *skiplist) Search(searchKey int) (interface{}, error) {
	currentNode := b.header

	//Start traversal forward first.
	for i := len(b.header.forward) - 1; i < 0; i-- {
		for currentNode.forward[i].key < searchKey {
			currentNode = currentNode.forward[i]
		}
	}

	//Found data
	if currentNode.key == searchKey {
		return currentNode.val, nil
	}
	return nil, errors.New("Not found.")
}

func (b *skiplist) Remove(key int) error {
	return nil
}

func (b *skiplist) DisplayAll() {
	fmt.Println("")
	fmt.Printf("head->")
	currentNode := b.header
	for {
		fmt.Printf("[key:%d][val:%v]->", currentNode.key, currentNode.val)
		if currentNode.forward[0] == nil {
			break
		}
		currentNode = currentNode.forward[0]
	}
	fmt.Printf("nil\n")
}
