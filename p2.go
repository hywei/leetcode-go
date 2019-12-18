package p2

import (
	"fmt"
	"math/big")

type ListNode struct {
	Val int
	Next *ListNode
}

func (L *ListNode) Append(v int) {
    list := &ListNode{
        Next: nil,
        Val:  v,
    }
	
	for k := L; k != nil; k = k.Next {
		if k.Next == nil {
			k.Next = list
			return 
		}
	}
	
}

func (list *ListNode) String() string {
	
	var s string
	if list != nil {
		s = fmt.Sprintf("%d", list.Val)
	}
	for k := list.Next; k != nil; k = k.Next {
		s = fmt.Sprintf("%s->%d", s, k.Val) 
	}
	
	return s
	
}

func toString(list *ListNode) string {
	var s string
	for k := list; k != nil; k = k.Next {
		s = string(int('0') + k.Val) + s
	}

	return s
}

func addString(s1 string, s2 string) string {
	k1, k2, k3 := new(big.Int), new(big.Int), new(big.Int)
	k1.SetString(s1, 10)
	k2.SetString(s2, 10)
	
	k3.Add(k1, k2)
	
	return fmt.Sprintf("%s", k3)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1, s2 := toString(l1), toString(l2)
	s := addString(s1, s2)
	fmt.Println(s1, s2, s, len(s), string(s[len(s)-1]))
	
	if addString(s1, s2) == "0" {
        return &ListNode { Next: nil, Val: 0}
    }
    
	var head, tail *ListNode
	for s := addString(s1, s2); len(s) != 0; s = s[:len(s)-1] {
		k := &ListNode{ Next: nil, Val: int(s[len(s)-1]) - int('0') }
		if head == nil {
			head, tail = k, k
		} else {
			tail.Next = k
			tail = k
		}
	}
	
	return head
	
}
