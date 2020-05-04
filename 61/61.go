package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	h := head
	len := 1
	for {
		if h.Next == nil {
			h.Next = head
			break
		}

		h = h.Next
		len++
	}
	r := h
	h = head
	k = k % len
	for i := 1; i <= len-k; i++ {
		r = r.Next
		h = h.Next
	}
	r.Next = nil
	return h
}

