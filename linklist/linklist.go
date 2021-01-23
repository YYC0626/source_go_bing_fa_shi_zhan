package linklist

import (
	"fmt"
)

type LinkList struct {
	head   *ListNode
	length uint
}

func CreateLinkList() *LinkList {
	return &LinkList{
		head:   &ListNode{0, nil},
		length: 0,
	}
}

func (this *LinkList) PrintLink() {
	cur := this.head.GetNext()

	for nil != cur {
		fmt.Printf("%v->", cur.GetValue())
		cur = cur.GetNext()
	}
	fmt.Println()
}
