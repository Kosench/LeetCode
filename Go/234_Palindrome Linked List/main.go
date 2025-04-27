package _34_Palindrome_Linked_List

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	// Input: head = [1,2,2,1]
	lastNode := reverseListNode(middleNode(head))
	// head = [1,2,2,nil]
	// lastNode = [1,2, nil]

	for lastNode != nil {
		if head.Val != lastNode.Val {
			return false
		}
		head = head.Next
		lastNode = lastNode.Next
	}
	return true
}

func reverseListNode(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		tmp := curr
		curr = curr.Next
		tmp.Next = prev
		prev = tmp
	}

	return prev
}

func middleNode(head *ListNode) *ListNode {
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}
