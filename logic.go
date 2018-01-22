package Answer

import (
	"sort"
)

//Answers 类型
type Answers struct {
	content string
	sum     int32
	freq    int32
}

//AnswersList 类型
type AnswersList []*Answers

func (l AnswersList) Len() int {
	return len(l)
}

func (l AnswersList) Less(i, j int) bool {
	if l[i].freq < l[j].freq {
		return true
	} else if l[i].freq > l[j].freq {
		return false
	} else {
		return l[i].sum < l[j].sum
	}
}

func (l AnswersList) Swap(i, j int) {
	temp := l[i]
	l[i] = l[j]
	l[j] = temp
}

func processLogic(input []*Answers, flag bool) (result string) {
	p := AnswersList(input)
	sort.Sort(p)
	if flag {
		result = p[0].content
	} else {
		result = p[len(p)-1].content
	}
	return result
}
