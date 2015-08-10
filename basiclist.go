package skiplist

import (
	"errors"
	"fmt"
)

type basicnote struct {
	key      int
	val      interface{}
	nextNode *basicnote
}

type basiclist struct {
	head *basicnote
}

//NewBasicList : Init structure for basic Sorted Linked List.
func NewBasicList() *basiclist {
	var empty interface{}
	return &basiclist{head: &basicnote{key: 0, val: empty, nextNode: nil}}
}

func (b *basiclist) Insert(key int, value interface{}) {
	if b.head == nil {
		// fmt.Println("note is empty")
		b.head = &basicnote{key: key, val: value, nextNode: nil}
	} else {
		var currentNode *basicnote
		currentNode = b.head
		var previouNote *basicnote
		var found bool
		newNode := &basicnote{key: key, val: value, nextNode: nil}

		for {
			if currentNode.key > key {
				newNode.nextNode = previouNote.nextNode
				previouNote.nextNode = newNode
				found = true
				break
			}

			if currentNode.nextNode == nil {
				break
			}

			previouNote = currentNode
			currentNode = currentNode.nextNode
		}

		if found == false {
			currentNode.nextNode = newNode
		}
	}
}

func (b *basiclist) Search(key int) (interface{}, error) {
	currentNode := b.head
	for {
		if currentNode.key == key {
			return currentNode.val, nil
		}

		if currentNode.nextNode == nil {
			break
		}
		currentNode = currentNode.nextNode
	}
	return nil, errors.New("Not found.")
}

func (b *basiclist) Remove(key int) error {
	currentNode := b.head
	var previouNote *basicnote
	for {
		if currentNode.key == key {
			previouNote.nextNode = currentNode.nextNode
			return nil
		}

		if currentNode.nextNode == nil {
			break
		}
		previouNote = currentNode
		currentNode = currentNode.nextNode
	}
	return errors.New("Not found key.")
}

func (b *basiclist) DisplayAll() {
	fmt.Println("")
	fmt.Printf("head->")
	currentNode := b.head
	for {
		fmt.Printf("[key:%d][val:%v]->", currentNode.key, currentNode.val)
		if currentNode.nextNode == nil {
			break
		}
		currentNode = currentNode.nextNode
	}
	fmt.Printf("nil\n")
}
