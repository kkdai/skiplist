package skiplist_test

import (
	"testing"

	. "github.com/kkdai/skiplist"
)

func TestRandomMaxLevel(t *testing.T) {
	sp := NewSkipList()
	sp.SetMaxLevel(5)
	for i := 0; i <= 5; i++ {
		ranLevel := sp.RandomLevel()
		if ranLevel > 5 || ranLevel < 0 {
			t.Errorf(" error found in random: %d", ranLevel)
		}
	}
}

func TestInsert(t *testing.T) {
	sl := NewSkipList()
	sl.Insert(30, "string 30")
	sl.Insert(50, "string 50")
	sl.Insert(40, "string 40")
	sl.Insert(20, "string 20")

	if val, _ := sl.Search(20); val != "string 20" {
		t.Errorf("Expect 20, got %d \n", val)
	}
}

func TestInsertBig(t *testing.T) {
	sl := NewSkipList()

	var i uint32
	for i = 0; i < 100000; i++ {
		sl.Insert(i, i)
	}

	val, _ := sl.Search(65536)
	ret := val.(uint32)
	if ret != 65536 {
		t.Error("Cannot search middle ret:", ret, val)
	}
}

func TestSearch(t *testing.T) {
	sl := NewSkipList()
	sl.Insert(30, "string 30")
	sl.Insert(50, "string 50")
	sl.Insert(40, "string 40")

	val, err := sl.Search(40)
	if err != nil || val != "string 40" {
		t.Errorf("search error, expect string40\n")
	}
}

func TestDelete(t *testing.T) {
	sl := NewSkipList()
	sl.Insert(30, "string 30")
	sl.Insert(50, "string 50")
	sl.Insert(40, "string 40")

	err := sl.Delete(40)
	if err != nil {
		t.Error("delete error")
	}
}

func TestDeleteBig(t *testing.T) {
	sl := NewSkipList()
	var i uint32
	for i = 0; i < 100000; i++ {
		sl.Insert(i, i)
	}

	for i = 0; i < 100000; i++ {
		err := sl.Delete(i)
		if err != nil {
			t.Error("delete big error")
		}
	}
}

/* Bigger than 25868 times delete might cause error
func BenchmarkDelete(b *testing.B) {
	sl := NewSkipList()
	var i uint32
	for i = 0; i < 100000; i++ {
		sl.Insert(i, i)
	}

	b.ResetTimer()
	fmt.Println(b.N)
	for i = 0; i < uint32(b.N); i++ {
		sl.Delete(i)
	}
}
*/

func BenchmarkSkiplistSearch(b *testing.B) {
	sl := NewSkipList()
	var i uint32
	for i = 0; i < 100000; i++ {
		sl.Insert(i, i)
	}

	b.ResetTimer()
	for i = 0; i < uint32(b.N); i++ {
		sl.Search(i)
	}
}

func BenchmarkSkiplistInsert(b *testing.B) {
	sl := NewSkipList()
	b.ResetTimer()
	var i uint32
	for i = 0; i < uint32(b.N); i++ {
		sl.Insert(i, i)
	}
}

func BenchmarkSliceInsert(b *testing.B) {
	var sl []uint32
	b.ResetTimer()
	var i uint32
	for i = 0; i < uint32(b.N); i++ {
		sl = append(sl, i)
	}
}

func BenchmarkSliceSearch(b *testing.B) {
	var sl []uint32
	var i uint32
	for i = 0; i < 100000; i++ {
		sl = append(sl, i)
	}

	var ret int
	b.ResetTimer()
	for i = 0; i < uint32(b.N); i++ {
		for k, v := range sl {
			if v == i {
				ret = k
				ret = ret + 1
			}
		}
	}
}
