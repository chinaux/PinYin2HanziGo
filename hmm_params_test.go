package PY2Hanzhi

import "testing"

func TestParamInit(t *testing.T) {
	hParam := HmmParams{}
	hParam.Init("./data")
}