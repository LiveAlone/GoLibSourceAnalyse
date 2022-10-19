package util

import (
	"fmt"
	"testing"

	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/sets/hashset"
)

func TestNone(t *testing.T) {
	list := arraylist.New()
	list.Add("a")      // ["a"]
	list.Add("c", "b") // ["a","c","b"]
	fmt.Println(list)
}

func TestSet(t *testing.T) {
	set := hashset.New()
	set.Add(123, 12, 345)
	set.Add(12, 34, 56)
	fmt.Println(set)

}
