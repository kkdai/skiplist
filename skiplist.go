package skiplist

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type skipnode struct {
	key     int
	val     interface{}
	forward []*skipnode
}

type skiplist struct {
	header     *skipnode
	maxLevel   int
	posibility float32
}

//NewSkipList : Init structure for basic Sorted Linked List.
func NewSkipList() *skiplist {
	var empty interface{}
	return &skiplist{header: &skipnode{key: 0, val: empty, forward: nil}}
}

func randomP() float32 {
	rand.Seed(int64(time.Now().Nanosecond()))
	return rand.Float32()
}

func (s *skiplist) randomLevel() int {
	level := 1
	for randomP() < s.posibility && level < s.maxLevel {
		level++
	}
	return level
}

//Search.
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

//Insert.
func (b *skiplist) Insert(searchKey int, value interface{}) {
	//TODO. local update MaxLevel

	currentNode := b.header

	for i := b.maxLevel; i <= 0; i-- {
		for currentNode.forward[i].key < searchKey {
			currentNode = currentNode.forward[i]
		}
		//TODO. update[i] = currentNode
	}

	currentNode = currentNode.forward[1] //TODO. Need checl why it is [1]
	if currentNode.key == searchKey {
		currentNode.val = value
	} else {
		levelV := b.randomLevel()
		if levelV > b.maxLevel {
			for i := 0; i <= b.maxLevel; i++ {
				//TODO. update[i] =
			}
		}
	}
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
