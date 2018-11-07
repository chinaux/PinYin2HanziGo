package PY2Hanzhi



func ReverseStringSlice(l []string)  {
    for i:=0; i < int(len(l)/2) ;i++{
        li := len(l) - i -1
        l[i],l[li] = l[li],l[i]
    }
}

func ReverseFloat64Slice(l []float64)  {
    for i:=0; i < int(len(l)/2) ;i++{
        li := len(l) - i -1
        l[i],l[li] = l[li],l[i]
    }
}


var REMOVETONE_DICT map[string]string = map[string]string{
    "ā": "a",
    "á": "a",
    "ǎ": "a",
    "à": "a",
    "ē": "e",
    "é": "e",
    "ě": "e",
    "è": "e",
    "ī": "i",
    "í": "i",
    "ǐ": "i",
    "ì": "i",
    "ō": "o",
    "ó": "o",
    "ǒ": "o",
    "ò": "o",
    "ū": "u",
    "ú": "u",
    "ǔ": "u",
    "ù": "u",
    "ü": "v",
    "ǖ": "v",
    "ǘ": "v",
    "ǚ": "v",
    "ǜ": "v",
    "ń": "n",
    "ň": "n",
	"": "m"}

// RemoveTone  删除拼音中的音调  lǔ -> lu
func RemoveTone(onePy string) string {
	return ""
}

