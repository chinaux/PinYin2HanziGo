package PY2Hanzhi


import(
	"strings"
	"math"
)

//Viterbi 
func Viterbi(hParams *HmmParams,observations string, 
	path_num int, log bool, min_prob float64) ([]float64,[]string) {
	
	v := []map[string]*PrioritySet{}
	v = append(v,make(map[string]*PrioritySet))	

	observationsList := strings.Split(observations," ")
	t := 0
    curObs := observationsList[t]	
	//Initialize base case (t == 0)
	curStatesList := hParams.GetState(curObs)	//wordset
	prevStatesList := curStatesList

	for _,state := range curStatesList {
		score := float64(0.0)
		if log {
			score = math.Log(math.Max(hParams.Start(state),min_prob)) + 
					math.Log(math.Max(hParams.Emission(state,curObs),min_prob))
		}else{
			score = math.Max(hParams.Start(state),min_prob) * 
					math.Max(hParams.Emission(state,curObs),min_prob)
		}
		pSet := &PrioritySet{Capacity:path_num,PathQueue:New()}
		pSet.Put(score,state)
		v[0][state] = pSet
	}

	// Run Viterbi for t > 0
    for t = 1; t < len(observationsList); t++ {
		curObs = observationsList[t]

		if len(v) == 2 {
			v = append(v[:0],v[1:]...) //删除元素 s = append(s[:i], s[i+1:]...)
		}
		v = append(v,make(map[string]*PrioritySet))	

		prevStatesList = curStatesList
		curStatesList = hParams.GetState(curObs)

		for _, y := range curStatesList {
			curSet := &PrioritySet{Capacity:path_num,PathQueue:New()}
			v[1][y] = curSet
			for _,y0 := range prevStatesList {
				prevSet,_ := v[0][y0]
				score := float64(0.0)
				for _, item := range prevSet.GetList() {
					if log {
						score = item.Score +
							math.Log(math.Max(hParams.Transition(y0,y),min_prob)) +
							math.Log(math.Max(hParams.Emission(y,curObs),min_prob))
					}else{
						
						score = item.Score *
							math.Max(hParams.Transition(y0,y),min_prob) *
							math.Max(hParams.Emission(y,curObs),min_prob)
					}
					p := item.Path + y
					v[1][y].Put(score,p)
				}
			}
		
		}

	}
       
	result := PrioritySet{Capacity:path_num,PathQueue:New()}
	lastNode := v[len(v)-1]
	for _,pSet := range lastNode {
		for _, item := range pSet.GetList() {
			result.Put(item.Score,item.Path)
		}
	}

	scoreList := []float64{}
	pathList := []string{}

	for result.PathQueue.Len() > 0 {
		x := result.PathQueue.Pop().(*Item)
		scoreList = append(scoreList,x.Score)
		pathList = append(pathList,x.Path)
	}

	ReverseStringSlice(pathList)
	ReverseFloat64Slice(scoreList)

	return scoreList,pathList

}