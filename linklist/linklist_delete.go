package linklist

import (
	"errors"
)

func (this *LinkList) Delete(node *ListNode) error {
	if nil == node {
		return errors.New("空节点")
	}
	pre := this.head
	cur := this.head.GetNext()

	for nil != cur {
		if cur.GetValue() == node.GetValue() {
			break
		}
		pre = cur
		cur = cur.GetNext()
	}
	if nil == cur {
		return errors.New("未找到节点")
	}
	pre.next = cur.GetNext()
	node = nil
	this.length--

	return nil
}
