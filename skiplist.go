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
	level   int
}

func NewNode(searchKey int, value interface{}, createLevel int, maxLevel int) *skipnode {
	//Every forward prepare a maxLevel empty point first.
	forwardEmpty := make([]*skipnode, maxLevel)
	for i := 0; i <= maxLevel-1; i++ {
		forwardEmpty[i] = nil
	}
	return &skipnode{key: searchKey, val: value, forward: forwardEmpty, level: createLevel}
}

type skiplist struct {
	header *skipnode
	//List configuration
	maxLevel    int
	propobility float32
	//List status
	level int //current level of whole skiplist
}

const (
	DefaultMaxLevel    int     = 4    //Maximal level allow to create in this skip list
	DefaultPropobility float32 = 0.50 //Default propobility
)

//NewSkipList : Init structure for basic Sorted Linked List.
func NewSkipList() *skiplist {
	newList := &skiplist{header: NewNode(0, "header", 1, DefaultMaxLevel), level: 1}
	newList.maxLevel = DefaultMaxLevel       //default
	newList.propobility = DefaultPropobility //default
	return newList
}

func randomP() float32 {
	rand.Seed(int64(time.Now().Nanosecond()))
	return rand.Float32()
}

func (s *skiplist) randomLevel() int {
	level := 1
	for randomP() < s.propobility && level < s.maxLevel {
		level++
	}
	return level
}

//Search.
func (b *skiplist) Search(searchKey int) (interface{}, error) {
	currentNode := b.header
	// fmt.Println("Head node level=", b.level)

	//Start traversal forward first.
	for i := b.level - 1; i >= 0; i-- {
		for currentNode.forward[i] != nil && currentNode.forward[i].key < searchKey {
			currentNode = currentNode.forward[i]
		}
	}

	//Step to final search node.
	currentNode = currentNode.forward[0]

	if currentNode.key == searchKey {
		return currentNode.val, nil
	}
	return nil, errors.New("Not found.")
}

//Insert.
func (b *skiplist) Insert(searchKey int, value interface{}) {
	updateList := make([]*skipnode, b.maxLevel)
	currentNode := b.header

	//Quick search in forward list
	for i := currentNode.level - 1; i >= 0; i-- {
		for currentNode.forward[i] != nil && currentNode.forward[i].key < searchKey {
			currentNode = currentNode.forward[i]
		}
		updateList[i] = currentNode
	}

	//Step to next node. (which is the target insert location)
	currentNode = currentNode.forward[0]

	if currentNode != nil && currentNode.key == searchKey {
		currentNode.val = value
	} else {
		newLevel := b.randomLevel()
		if newLevel > b.level {
			for i := b.level + 1; i <= newLevel; i++ {
				updateList[i-1] = b.header
			}
			b.level = newLevel //This is not mention is pseudo code
			b.header.level = newLevel
		}

		newNode := NewNode(searchKey, value, newLevel, b.maxLevel) //New node
		for i := 0; i <= newLevel-1; i++ {                         //zero base
			newNode.forward[i] = updateList[i].forward[i]
			updateList[i].forward[i] = newNode
		}
	}
}

func (b *skiplist) Delete(searchKey int) error {
	updateList := make([]*skipnode, b.maxLevel)
	currentNode := b.header

	//Quick search in forward list
	for i := currentNode.level - 1; i >= 0; i-- {
		for currentNode.forward[i] != nil && currentNode.forward[i].key < searchKey {
			currentNode = currentNode.forward[i]
		}
		updateList[i] = currentNode
	}

	//Step to next node. (which is the target delete location)
	currentNode = currentNode.forward[0]

	if currentNode.key == searchKey {
		for i := 0; i <= currentNode.level-1; i++ {
			if updateList[i].forward[i] != nil && updateList[i].forward[i].key != currentNode.key {
				break
			}
			updateList[i].forward[i] = currentNode.forward[i]
		}

		for currentNode.level > 1 && b.header.forward[currentNode.level] == nil {
			currentNode.level--
		}

		//free(currentNode)  //no need for Golang because GC
		currentNode = nil
		return nil
	}
	return errors.New("Not found")
}

func (b *skiplist) DisplayAll() {
	fmt.Printf("\nhead->")
	currentNode := b.header

	//Draw forward[0] base
	for {
		fmt.Printf("[key:%d][val:%v]->", currentNode.key, currentNode.val)
		if currentNode.forward[0] == nil {
			break
		}
		currentNode = currentNode.forward[0]
	}
	fmt.Printf("nil\n")

	fmt.Println("---------------------------------------------------------")
	currentNode = b.header
	//Draw all data node.
	for {
		fmt.Printf("[node:%d], val:%v, level:%d ", currentNode.key, currentNode.val, currentNode.level)

		if currentNode.forward[0] == nil {
			break
		}

		for j := currentNode.level - 1; j >= 0; j-- {
			fmt.Printf(" fw[%d]:", j)
			if currentNode.forward[j] != nil {
				fmt.Printf("%d", currentNode.forward[j].key)
			} else {
				fmt.Printf("nil")
			}
		}
		fmt.Printf("\n")
		currentNode = currentNode.forward[0]
	}
	fmt.Printf("\n")
}
