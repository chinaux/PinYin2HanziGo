package PY2Hanzhi

import(
	"encoding/json"
	"io/ioutil"
	"fmt"
	"strings"
)


type Emission struct {
	Data 	map[string]map[string]float64	// {"万":{"mo":1.05.346100491351276e-06,"wan":0.9999946538995087}}
	Default float64
}


type Transition struct {
	Data 	map[string]map[string]float64	// {"一":{"一":0.002499256744160108,"丈":2.2564782682252634e-05}}
	Default float64
}

type Start struct {
	Data	map[string]float64	//	"你":0.00027823067236989926
	Default float64
}

type HmmParams struct {
	PY2HZDict map[string]string		// "ni":"你尼泥..."
	StartDict 		Start
	EmissionDict	Emission	
	TransitionDict 	Transition
}
//Init 初始化加载data目录中的各概率矩阵
func (hParam *HmmParams) Init(dataDir string) {
	hParam.PY2HZDict = make(map[string]string)

	LoadJsonFile(dataDir+"/hmm_py2hz.json",&hParam.PY2HZDict)
	LoadJsonFile(dataDir+"/hmm_start.json",&hParam.StartDict)
	LoadJsonFile(dataDir+"/hmm_emission.json",&hParam.EmissionDict)
	LoadJsonFile(dataDir+"/hmm_transition.json",&hParam.TransitionDict)
	/*
	fmt.Println("hParam.PY2HZDict",len(hParam.PY2HZDict))
	fmt.Println("hParam.StartDict.Data",len(hParam.StartDict.Data))
	fmt.Println("hParam.StartDict.Default",hParam.StartDict.Default)
	fmt.Println("hParam.EmissionDict.Data",len(hParam.EmissionDict.Data))
	fmt.Println("hParam.EmissionDict.Default",hParam.EmissionDict.Default)
	fmt.Println("hParam.TransitionDict.Data",len(hParam.TransitionDict.Data))
	fmt.Println("hParam.TransitionDict.Default",hParam.TransitionDict.Default)
	*/
}

func (hParam *HmmParams) Start(state string) float64{
	if value,exist := hParam.StartDict.Data[state]; exist {
		return value
	}
	return hParam.StartDict.Default
}

func (hParam *HmmParams) Emission(state,observation string) float64{

	pinyin := observation
	hanzi := state

	if probDict,exist := hParam.EmissionDict.Data[hanzi]; exist {
		if prob,exist := probDict[pinyin]; exist {
			return prob
		}
	}

	return hParam.EmissionDict.Default
}


func (hParam *HmmParams) Transition(fromState,toState string) float64 {

	if probDict,exist := hParam.TransitionDict.Data[fromState]; exist {
		if prob,exist := probDict[toState]; exist {
			return prob
		}else{
			if prob,exist := probDict["default"]; exist {
				return prob
			}
		}
	}
	return hParam.TransitionDict.Default
}

func (hParam *HmmParams) GetState(observation string) []string {

	if wordStr,exist := hParam.PY2HZDict[observation]; exist {
		return strings.Split(wordStr,"")
	}
	return []string{}
}

func LoadJsonFile(filepath string,v interface{}) error {

	 //ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	 data, err := ioutil.ReadFile(filepath)
	 if err != nil {
		 fmt.Errorf("ReadFile err:",err)
		 return err
	 }
 
	 //读取的数据为json格式，需要进行解码

	 err = json.Unmarshal(data, &v)
	 if err != nil {
		fmt.Errorf("Unmarshal err:",err)
		 return err
	 }
	 return nil
}

