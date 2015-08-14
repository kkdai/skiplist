package skiplist

import "testing"

func TestRandomMaxLevel(t *testing.T) {
	sp := NewSkipList()
	sp.maxLevel = 5
	for i := 0; i <= sp.maxLevel; i++ {
		ranLevel := sp.randomLevel()
		if ranLevel > sp.maxLevel || ranLevel < 0 {
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

	if sl.header.forward[0].key != 20 {
		t.Errorf("Expect 20, got %d \n", sl.header.forward[0].key)
	}
}
