package MergeTwoSortedLists

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(list *ListNode) {
	p := list
	for p.Next != nil {
		fmt.Printf("p.value : %d ", p.Val)
		fmt.Println("p.next : ", p.Next)
		p = p.Next
	}
	fmt.Printf("p.value : %d ", p.Val)
	fmt.Println("p.next : ", p.Next)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	head := ListNode{}
	pointer := &head

	if l1 == nil && l2 == nil {
		return nil
	}

	for l1 != nil || l2 != nil {

		if l1 == nil {
			(*pointer) = *l2
			break
		} else if l2 == nil {
			(*pointer) = *l1
			break
		}

		if l1.Val <= l2.Val {
			(*pointer).Val = l1.Val
			l1 = l1.Next
		} else if l1.Val > l2.Val {
			(*pointer).Val = l2.Val
			l2 = l2.Next
		}
		(*pointer).Next = &ListNode{}
		pointer = (*pointer).Next

	}

	return &head
}
