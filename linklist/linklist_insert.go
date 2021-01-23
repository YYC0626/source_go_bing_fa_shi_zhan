package linklist

import (
	"errors"
)

func (this *LinkList) Insert(i uint, node *ListNode) error {

	if i < 0 || node == nil || i > this.length {
		return errors.New("空或越界")
	}

	curNode := (*this).head
	for j := uint(0); j < i; j++ {
		curNode = curNode.GetNext()
	}
	node.next = curNode.GetNext()
	curNode.next = node
	this.length++

	return nil
}
