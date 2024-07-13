package gmap_test

import (
	"fmt"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	condition := []string{"a=1", "b=2"}
	//var conditionList []string
	//for k, v := range condition {
	//	cond := fmt.Sprintf(`%s="%s"`, k, v)
	//	fmt.Println(k, v, cond)
	//	conditionList = append(conditionList, cond)
	//}
	//fmt.Println(len(conditionList))
	//fmt.Println(strings.Join(conditionList,  " and "))

	fmt.Println(strings.Join(condition, " and "))
}
