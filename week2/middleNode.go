package main

type ListNode struct {
	     Val int
	    Next *ListNode
	 }

func middleNode(head *ListNode) *ListNode {

	node:= head;
	fast:= head;

	for fast != nil && fast.Next != nil {
		node = node.Next;
		fast = fast.Next.Next;
	}

	return node;
}

// func middleNode(head *ListNode) *ListNode {

// 	node:= head;
// 	fast:= head;

// 	c := 0;

// 	for fast != nil {
// 		fast = fast.Next
// 	}
// 	index:= (c/ 2);
// 	if c % 2 > 0{
// 		index++;
// 	}

// 	node = head
// 	for true {
// 		if index == 0{
// 			return node;
// 		}

// 		node = node.Next;
// 		index--;
// 	}

// 	return node;
// }