package PY2Hanzhi

import (
	"fmt"
	"testing"
)


func TestViterbi(t *testing.T) {

	hmmparams := &HmmParams{}
	hmmparams.Init("./data")

	scoreList,pathList := Viterbi(hmmparams, "ni zhi bu zhi dao",2,false,3.14e-200)

	fmt.Println(scoreList)
	fmt.Println(pathList)
}