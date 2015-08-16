package skiplist_test

import (
	"fmt"

	. "github.com/kkdai/skiplist"
)

func Example_manuplateSkiplist() {
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

	//Delete by search key
	err = sl.Delete(70)
	if err != nil {
		fmt.Println("Delete not found")
	}

	//Display all skip list content.
	sl.DisplayAll()

}
