package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//Recursive Solution
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    } else {
        reverseLinkedList := reverseList(head.Next)
        head.Next.Next = head //Next nodes next pointer is set to current node
        head.Next = nil // Current node next pointer is set to nil
        //Recursively this means only the tail has a next pointer of nil
        return reverseLinkedList
    }
}

//Iterative Solution
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    curr := head
    for curr != nil {
        tmp := curr.Next
        curr.Next = prev
        prev = curr
        curr = tmp
    }
    return prev
}
