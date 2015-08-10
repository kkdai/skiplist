package skiplist

import (
	"fmt"
	"log"
)

type basicnote struct {
	key      int
	val      interface{}
	nextNode *basicnote
}

type basiclist struct {
	head *basicnote
}

func NewBasicList() *basiclist {
	var empty interface{}
	return &basiclist{head: &basicnote{key: 0, val: empty, nextNode: nil}}
}

func (b *basiclist) Insert(key int, value interface{}) {
	if b.head == nil {
		log.Println("note is empty")
		b.head = &basicnote{key: key, val: value, nextNode: nil}
	} else {
		var currentNode *basicnote
		currentNode = b.head
		var previouNote *basicnote
		var found bool
		newNode := &basicnote{key: key, val: value, nextNode: nil}

		log.Println("currentNode:", currentNode)
		for currentNode.nextNode != nil {
			if currentNode.key > key {
				newNode.nextNode = previouNote.nextNode
				previouNote.nextNode = newNode
				log.Println("insert node currentNode:", currentNode)
				found = true
				return
			}
			previouNote = currentNode
			currentNode = b.head.nextNode
			log.Println("currentNode:", currentNode)
		}

		if found == false {
			log.Println("Not found, assign to latest")
			b.head.nextNode = newNode
		}
	}
}

func (b *basiclist) DisplayAll() {
	fmt.Println("")
	fmt.Printf("head->")
	currentNode := b.head
	for currentNode.nextNode != nil {
		fmt.Printf("[key:%d][val:%v]->", currentNode.key, currentNode.val)
		currentNode = currentNode.nextNode
	}
	fmt.Printf("nil\n")
}
