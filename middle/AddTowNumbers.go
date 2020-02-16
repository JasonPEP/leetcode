package middle

/**
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.

*/
// the function run result is pass, but the performance is poor
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	m := make(map[int]*ListNode)
	for i, carry := 0, 0; ; i++ {
		sum := carry
		carry = 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		if sum >= 10 {
			carry = 1
			sum %= 10
		}
		m[i] = &ListNode{
			Val: sum,
		}
		if l1 == nil && l2 == nil && carry == 0 {
			break
		}
	}

	for k := range m {
		if k > 0 {
			m[k-1].Next = m[k]
		}
	}
	return m[0]
}

// the function do not cover big numbers
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	l1Num := 0
	l2Num := 0
	for pop := 1; l1 != nil || l2 != nil; pop = pop * 10 {
		if l1 != nil {
			l1Num += l1.Val * pop
			l1 = l1.Next
		}
		if l2 != nil {
			l2Num += l2.Val * pop
			l2 = l2.Next
		}
	}
	l1, l2 = nil, nil
	var head = &ListNode{
		Next: &ListNode{},
	}
	var cur = head
	sum := l1Num + l2Num
	for sum > 0 {
		cur.Next = &ListNode{
			Val: sum % 10,
		}
		cur = cur.Next
		sum = sum / 10
	}
	return head.Next
}

// the function run result is pass
// run cases 1563, running time 8ms
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	temp := result
	for i, carry := 0, 0; ; i++ {
		sum := carry
		carry = 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		if sum >= 10 {
			carry = 1
			sum %= 10
		}
		temp.Next = &ListNode{
			Val: sum,
		}
		temp = temp.Next
		if l1 == nil && l2 == nil && carry == 0 {
			break
		}
	}
	return result.Next
}

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}
