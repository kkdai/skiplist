package main

import (
	"fmt"

	skiplist "../../go-skiplist"
)

func main() {
	sl := skiplist.NewSkipList()
	// sl.Insert(30, "string 1")
	sl.Insert(50, "5")
	sl.Insert(40, "4")
	sl.Insert(70, "7")
	sl.Insert(100, "10")
	sl.Insert(10, "1")
	sl.Insert(20, "2")
	sl.Insert(30, "3")
	sl.Insert(80, "8")
	sl.Insert(90, "9")
	sl.DisplayAll()

	ret, err := sl.Search(70)
	if err == nil {
		fmt.Println("70->", ret)
	} else {
		fmt.Println("Not found, ", err)
	}
}
