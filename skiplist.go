package skiplist

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Skipnode struct {
	Key     int
	Val     interface{}
	Forward []*Skipnode
	Level   int
}

func NewNode(searchKey int, value interface{}, createLevel int, maxLevel int) *Skipnode {
	//Every forward prepare a maxLevel empty point first.
	forwardEmpty := make([]*Skipnode, maxLevel)
	for i := 0; i <= maxLevel-1; i++ {
		forwardEmpty[i] = nil
	}
	return &Skipnode{Key: searchKey, Val: value, Forward: forwardEmpty, Level: createLevel}
}

type Skiplist struct {
	Header *Skipnode
	//List configuration
	MaxLevel    int
	Propobility float32
	//List status
	Level int //current level of whole skiplist
}

const (
	DefaultMaxLevel    int     = 4    //Maximal level allow to create in this skip list
	DefaultPropobility float32 = 0.50 //Default propobility
)

//NewSkipList : Init structure for Skit List.
func NewSkipList() *Skiplist {
	newList := &Skiplist{Header: NewNode(0, "header", 1, DefaultMaxLevel), Level: 1}
	newList.MaxLevel = DefaultMaxLevel       //default
	newList.Propobility = DefaultPropobility //default
	return newList
}

func randomP() float32 {
	rand.Seed(int64(time.Now().Nanosecond()))
	return rand.Float32()
}

//Change SkipList default maxlevel is 4.
func (b *Skiplist) SetMaxLevel(maxLevel int) {
	b.MaxLevel = maxLevel
}

func (b *Skiplist) RandomLevel() int {
	level := 1
	for randomP() < b.Propobility && level < b.MaxLevel {
		level++
	}
	return level
}

//Search: Search a element by search key and return the interface{}
func (b *Skiplist) Search(searchKey int) (interface{}, error) {
	currentNode := b.Header

	//Start traversal forward first.
	for i := b.Level - 1; i >= 0; i-- {
		for currentNode.Forward[i] != nil && currentNode.Forward[i].Key < searchKey {
			currentNode = currentNode.Forward[i]
		}
	}

	//Step to final search node.
	currentNode = currentNode.Forward[0]

	if currentNode.Key == searchKey {
		return currentNode.Val, nil
	}
	return nil, errors.New("Not found.")
}

//Insert: Insert a search key and its value which could be interface.
func (b *Skiplist) Insert(searchKey int, value interface{}) {
	updateList := make([]*Skipnode, b.MaxLevel)
	currentNode := b.Header

	//Quick search in forward list
	for i := b.Header.Level - 1; i >= 0; i-- {
		for currentNode.Forward[i] != nil && currentNode.Forward[i].Key < searchKey {
			currentNode = currentNode.Forward[i]
		}
		updateList[i] = currentNode
	}

	//Step to next node. (which is the target insert location)
	currentNode = currentNode.Forward[0]

	if currentNode != nil && currentNode.Key == searchKey {
		currentNode.Val = value
	} else {
		newLevel := b.RandomLevel()
		if newLevel > b.Level {
			for i := b.Level + 1; i <= newLevel; i++ {
				updateList[i-1] = b.Header
			}
			b.Level = newLevel //This is not mention in cookbook pseudo code
			b.Header.Level = newLevel
		}

		newNode := NewNode(searchKey, value, newLevel, b.MaxLevel) //New node
		for i := 0; i <= newLevel-1; i++ {                         //zero base
			newNode.Forward[i] = updateList[i].Forward[i]
			updateList[i].Forward[i] = newNode
		}
	}
}

//Delete: Delete element by search key
func (b *Skiplist) Delete(searchKey int) error {
	updateList := make([]*Skipnode, b.MaxLevel)
	currentNode := b.Header

	//Quick search in forward list
	for i := b.Header.Level - 1; i >= 0; i-- {
		for currentNode.Forward[i] != nil && currentNode.Forward[i].Key < searchKey {
			currentNode = currentNode.Forward[i]
		}
		updateList[i] = currentNode
	}

	//Step to next node. (which is the target delete location)
	currentNode = currentNode.Forward[0]

	if currentNode.Key == searchKey {
		for i := 0; i <= currentNode.Level-1; i++ {
			if updateList[i].Forward[i] != nil && updateList[i].Forward[i].Key != currentNode.Key {
				break
			}
			updateList[i].Forward[i] = currentNode.Forward[i]
		}

		for currentNode.Level > 1 && b.Header.Forward[currentNode.Level] == nil {
			currentNode.Level--
		}

		//free(currentNode)  //no need for Golang because GC
		currentNode = nil
		return nil
	}
	return errors.New("Not found")
}

//DisplayAll: Display current SkipList content in console, will also print out the linked pointer.
func (b *Skiplist) DisplayAll() {
	fmt.Printf("\nhead->")
	currentNode := b.Header

	//Draw forward[0] base
	for {
		fmt.Printf("[key:%d][val:%v]->", currentNode.Key, currentNode.Val)
		if currentNode.Forward[0] == nil {
			break
		}
		currentNode = currentNode.Forward[0]
	}
	fmt.Printf("nil\n")

	fmt.Println("---------------------------------------------------------")
	currentNode = b.Header
	//Draw all data node.
	for {
		fmt.Printf("[node:%d], val:%v, level:%d ", currentNode.Key, currentNode.Val, currentNode.Level)

		if currentNode.Forward[0] == nil {
			break
		}

		for j := currentNode.Level - 1; j >= 0; j-- {
			fmt.Printf(" fw[%d]:", j)
			if currentNode.Forward[j] != nil {
				fmt.Printf("%d", currentNode.Forward[j].Key)
			} else {
				fmt.Printf("nil")
			}
		}
		fmt.Printf("\n")
		currentNode = currentNode.Forward[0]
	}
	fmt.Printf("\n")
}
