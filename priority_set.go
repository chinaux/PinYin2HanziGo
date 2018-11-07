package PY2Hanzhi


type PrioritySet struct {
	Capacity 	int
	PathQueue	*PriorityQueue
}

type Item struct {
	Score float64
	Path  string
}

func (this *Item) Less(other interface{}) bool {
	return this.Score < other.(*Item).Score
}


func (pSet *PrioritySet) Put(score float64,state string) {
	pSet.PathQueue.Push(&Item{Score:score,Path:state})
	if pSet.PathQueue.Len() > pSet.Capacity {
		pSet.PathQueue.Pop()
	}
}


func (pSet *PrioritySet) GetList() []*Item {
	ret := []*Item{}
	for _,it := range *pSet.PathQueue.s {
		ret = append(ret,it.(*Item))
	}
	return ret
}