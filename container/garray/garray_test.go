package garray_test

import (
	"fmt"
	"github.com/ilylx/gconv/container/garray"
	"testing"
)

func TestArray_Contains(t *testing.T) {
	ids := []int{1, 2, 3}
	var id int
	id = 1
	fmt.Println(garray.NewIntArrayFrom(ids).Contains(id))

}

func TestArrayUin64_Contains(t *testing.T) {
	ids := []uint64{1123123123123213, 123123123123213, 12321312312387}

	fmt.Println(garray.NewUint64From(ids).Contains(123213123123871))

}
