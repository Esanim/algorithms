package datastructure

import (
	"strconv"
	"testing"

	"github.com/maximelamure/algorithms/common"
)

func TestDoubleLinkedList(t *testing.T) {
	helper := common.Test{}
	list := &list{}

	list.Insert(5)
	list.Insert(6)
	list.Insert(7)
	len := list.Length()
	helper.Assert(t, len == 3, "The length should be 3. Here: "+strconv.Itoa(len))
	last := list.Last()
	helper.Assert(t, last.Val == 5, "The value of the last node should be 5. Here: "+strconv.Itoa(last.Val.(int)))
	list.Remove(6)
	len = list.Length()
	helper.Assert(t, len == 2, "After removing 1 element, the length should be 2. Here: "+strconv.Itoa(len))
}
