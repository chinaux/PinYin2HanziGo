package PY2Hanzhi

import (
	"fmt"
	"testing"
)

func TestPrioritySet(t *testing.T) {
	pSet := PrioritySet{Capacity:6,PathQueue:New()}
	pSet.Put(1.0,"你")
	pSet.Put(0.5,"好")
	pSet.Put(1.5,"吗")
	pSet.Put(2.0,"呀")
	itemList := pSet.GetList()
	for _,item := range itemList {
		fmt.Println(item.Score,item.Path)
	}

	for pSet.PathQueue.Len() > 0 {
		x := pSet.PathQueue.Pop().(*Item)
		fmt.Println(x.Score,x.Path)
	}
}