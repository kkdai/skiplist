package skiplist_test

import (
	"fmt"

	. "github.com/kkdai/skiplist"
)

func ExampleNewSkipList() {
	//New a skiplist
	sl := NewSkipList()
	sl.DisplayAll()
}

func ExampleSkiplist_Insert() {
	//New a skiplist
	sl := NewSkipList()

	//Insert search key 50, value "5", value could be anything.
	sl.Insert(50, "5")
}

func ExampleSkiplist_Search() {
	//New a skiplist
	sl := NewSkipList()

	//Insert search key 50, value "5", value could be anything.
	sl.Insert(50, "5")
	sl.Insert(40, "4")
	sl.Insert(70, "7")
	sl.Insert(100, "10")

	//Search key, which time complexity O(log n)
	ret, err := sl.Search(50)
	if err == nil {
		fmt.Println("key 50: val->", ret)
	} else {
		fmt.Println("Not found, ", err)
	}
}

func ExampleSkiplist_Delete() {
	//New a skiplist
	sl := NewSkipList()

	//Insert search key 50, value "5", value could be anything.
	sl.Insert(70, "7")

	//Delete by search key
	err := sl.Delete(70)
	if err != nil {
		fmt.Println("Delete not found")
	}
}
