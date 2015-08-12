package skiplist

import (
	"fmt"
	"testing"
	"time"
)

func TestRandomPossibility(t *testing.T) {
	fmt.Println(randomP())
	time.Sleep(time.Millisecond)
	fmt.Println(randomP())
	time.Sleep(time.Millisecond)
	fmt.Println(randomP())
}
